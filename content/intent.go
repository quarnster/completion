package content

import (
	"errors"
	"fmt"
	"io"
)

var handlers = map[string]Handler{}

type Handler func(io.Writer, Intent)

type Intent struct {
	Version   int64
	Operation string
}

type Response struct {
	Version int64 `json:"version"`
}

func AddHandler(name string, h Handler) {
	if _, ok := handlers[name]; ok {
		// immediately exit to prevent double-register which should
		// come up during development & testing.
		panic(fmt.Sprintf("Handler already registered for %s", name))
	}
	handlers[name] = h
}

func Handle(conn io.Writer, intent Intent) error {
	if handler, ok := handlers[intent.Operation]; !ok {
		return errors.New(fmt.Sprintf("Failed to locate handler for operation %s", intent.Operation))
	} else {
		handler(conn, intent)
	}
	return nil
}
