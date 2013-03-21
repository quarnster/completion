package content

import (
	"errors"
	"fmt"
	"io"
)

var handlers = map[string]Handler{}

type (
	IntentHandler interface {
		CanHandle(it *Intent) bool
		Handle(it *Intent) *Response
	}
	// TODO(@dskinner): Replace Handler with IntentHandler or something else?
	Handler func(io.Writer, Intent)

	Intent struct {
		Version   int64
		Operation string
		Data      Settings
	}

	Response struct {
		Version int64    `json:"version"`
		Data    Settings `json:"data"`
	}

	Session struct {
		Settings Settings
	}
)

func NewIntent(key string) (ret Intent) {
	ret.Operation = key
	ret.Data = *NewSettings()
	return
}

func NewResponse() (ret Response) {
	ret.Data = *NewSettings()
	return
}

// Returns the Session object associated with this Intent or nil.
func (it *Intent) Session() *Session {
	// TODO(someone): Intent.Session not implemented yet
	if sessionid, ok := it.Data.Get("sessionid").(int); ok {
		_ = sessionid // TODO actually look up session and return it
	}
	return nil
}

// If a Session is associated with this Intent, then the Session's
// Settings is cloned and the intent's Settings are merged into the
// cloned Settings object. This to allow Intents to be "slim", i.e.
// only contain the keys of the specific settings it wants to override.
//
// If there's no Session nor Intent Settings, the empty Settings object
// is returned.
func (it *Intent) Settings() *Settings {
	if session := it.Session(); session != nil {
		set := session.Settings.Clone()
		if settings, ok := it.Data.Get("settings").(Settings); ok {
			set.Merge(&settings)
		}
		return set
	} else if settings, ok := it.Data.Get("settings").(Settings); ok {
		return &settings
	}
	return NewSettings()
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
