// +build !nolibclang

package libclang

import (
	"bytes"
	"code.google.com/p/log4go"
	"fmt"
	"github.com/quarnster/completion/content"
	"github.com/sbinet/go-clang"
)

type Clang struct {
	TranslationUnitCache
}

func (t *Clang) CompleteAt(args *content.CompleteAtArgs, res *content.CompletionResult) error {
	origargs, _ := args.Settings().Get("compiler_flags").([]string)
	var unsaved map[string]string
	if args.Location.File.Contents != "" {
		unsaved = map[string]string{args.Location.File.Name: args.Location.File.Contents}
	}

	if tu := t.GetTranslationUnit(args.Location.File.Name, origargs, "", unsaved); tu == nil {
		return nil
	}
	cres := tu.CompleteAt(args.Location.File.Name, int(args.Location.Line), int(args.Location.Column), unsaved, 0)
	if !cres.IsValid() {
		return fmt.Errorf("CompleteResults is not valid")
	}
	defer cres.Dispose()
	for _, r := range cres.Results() {
		var (
			buf bytes.Buffer
		)
		switch r.CursorKind {
		case clang.CK_StructDecl, clang.CK_TypedefDecl:
			for _, c := range r.CompletionString.Chunks() {
				buf.WriteString(c.Text())
			}
			var tt content.Type
			tt.Flags = content.FLAG_TYPE_CLASS
			tt.Name.Absolute = buf.String()
			res.Types = append(res.Types, tt)
		case clang.CK_FunctionDecl:
			var (
				m            content.Method
				paramstarted bool
				argCount     int
			)
			for _, c := range r.CompletionString.Chunks() {
				switch k := c.Kind(); k {
				case clang.CompletionChunk_ResultType:
					var v content.Variable
					v.Name.Relative = c.Text()
					m.Returns = append(m.Returns, v)
				case clang.CompletionChunk_Placeholder:
					var v content.Variable
					v.Type.Name.Relative = c.Text()
					v.Name.Relative = fmt.Sprintf("arg%d", argCount)
					argCount++
					m.Parameters = append(m.Parameters, v)
				case clang.CompletionChunk_LeftParen:
					paramstarted = true
				case clang.CompletionChunk_RightParen, clang.CompletionChunk_Comma:
				case clang.CompletionChunk_TypedText:
					if !paramstarted {
						buf.WriteString(c.Text())
					}
				default:
					log4go.Warn("Unimplemented CompletionChunkKind: %s", k)
				}
			}
			m.Name.Relative = buf.String()
			res.Methods = append(res.Methods, m)
		default:
			log4go.Warn("Unimplemented CursorKind: %s", r.CursorKind)
		}
	}
	return nil
}
