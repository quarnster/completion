// +build !nolibclang

package libclang

import (
	"code.google.com/p/log4go"
	"github.com/quarnster/completion/util/expand_path"
	"github.com/sbinet/go-clang"
	"reflect"
	"sync"
)

type (
	LockedTranslationUnit struct {
		sync.Mutex
		clang.TranslationUnit
		opts        []string
		opts_script string
	}

	tulut map[string]*LockedTranslationUnit

	workunit struct {
		filename       string
		options        []string
		options_script string
		on_done        func()
	}

	workqueue chan func()

	TranslationUnitCache struct {
		index   clang.Index
		lut     tulut
		parsing map[string]bool
		queue   workqueue
		sync.Mutex
	}
)

const (
	WORK_QUEUE_LEN      = 10
	index_parse_options = clang.TU_PrecompiledPreamble | clang.TU_CacheCompletionResults
)

func NewTranslationUnitCache() *TranslationUnitCache {
	t := &TranslationUnitCache{}
	t.index = clang.NewIndex(0, 0)
	t.lut = make(tulut)
	t.parsing = make(map[string]bool)
	t.queue = make(workqueue, WORK_QUEUE_LEN)
	return t
}

func (t *TranslationUnitCache) parse(w workunit) (ret *LockedTranslationUnit) {
	// TODO(q): SublimeClang marked the file as "busy". Is that needed?
	defer func() {
		t.Lock()
		defer t.Unlock()
		t.parsing[w.filename] = false
		// TODO(q): Is on_done really needed?
		if w.on_done != nil {
			w.on_done()
		}
	}()
	log4go.Debug("Parsing %s", w.filename)
	ret = t.GetTranslationUnit(w.filename, w.options, w.options_script, nil)
	log4go.Debug("Parsing %s done", w.filename)
	return ret
}

func (t *TranslationUnitCache) reparse(w workunit) {
	if tu := t.parse(w); tu != nil {
		tu.Lock()
		defer tu.Unlock()

		log4go.Debug("Reparsing %s", w.filename)
		tu.Reparse(nil, 0)
		log4go.Debug("Reparsing %s done", w.filename)
	}
}

func (t *TranslationUnitCache) addEx(w workunit) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.lut[w.filename]; !ok && !t.parsing[w.filename] {
		t.parsing[w.filename] = true
		t.queue <- func() { t.parse(w) }
	}
}

func (t *TranslationUnitCache) GetTranslationUnit(filename string, options []string, options_script string, unsaved_files map[string]string) *LockedTranslationUnit {
	t.Lock()
	tu, ok := t.lut[filename]
	if !ok {
		t.Unlock()

		// TODO(q): SublimeClang executed opts_script here
		args := make([]string, len(options))
		for i := range options {
			args[i] = expand_path.ExpandPath(options[i])
		}

		log4go.Debug("Will compile file %s with the following options:\n%v", filename, options)
		if tu2 := t.index.Parse(filename, options, unsaved_files, index_parse_options); true || tu2.IsValid() {
			tu = &LockedTranslationUnit{}
			tu.TranslationUnit = tu2
			tu.opts = options
			tu.opts_script = options_script
			t.Lock()
			defer t.Unlock()
			t.lut[filename] = tu
		} else {
			log4go.Warn("Failed to compile %s, %v", filename, tu2)
		}
		return tu
	}
	recompile := !reflect.DeepEqual(tu.opts, options) || tu.opts_script != options_script
	if recompile {
		// TODO: need to dispose the tu.. Who's responsible for its disposal?
		delete(t.lut, filename)
	}
	t.Unlock()
	if recompile {
		log4go.Debug("Options change detected. Will recompile %s", filename)
		t.addEx(workunit{filename, options, options_script, nil})
	}
	return tu
}
