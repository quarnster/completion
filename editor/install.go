package editor

import (
	"archive/zip"
	"bytes"
	"code.google.com/p/log4go"
	"fmt"
	"io"
	"os"
)

type (
	EditorPlugin interface {
		Name() string
		Description() string
		Install() error
		Uninstall() error
	}
)

func Open(path string) (io.ReadCloser, error) {
	r, err := zip.NewReader(bytes.NewReader(archive_data), int64(len(archive_data)))
	if err != nil {
		return nil, err
	}
	var f1 io.ReadCloser
	for _, f := range r.File {
		if f.Name != path {
			continue
		}
		f1, err = f.Open()
		if err != nil {
			return nil, err
		}
		break
	}
	if f1 == nil {
		return nil, fmt.Errorf("File not found: %v", path)
	}
	return f1, nil
}

func Copy(a, b string) error {
	f1, err := Open(a)
	if err != nil {
		return err
	}
	defer f1.Close()
	f2, err := os.Create(b)
	if err != nil {
		return err
	}
	defer f2.Close()
	log4go.Debug("Copying %s to %s", a, f2.Name())
	if _, err := io.Copy(f2, f1); err != nil {
		return err
	}
	return nil
}

var plugins []EditorPlugin

func Register(plugin EditorPlugin) {
	plugins = append(plugins, plugin)
}

func List() []EditorPlugin {
	return plugins
}
