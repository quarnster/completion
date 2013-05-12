package content

import (
	"fmt"
)

type (
	Session struct {
		Settings
	}

	SessionRegisterArgs struct {
		SessionId string
		Settings  Settings
	}
)

var sessionmap = make(map[string]*Session)

// TODO(.): Not too happy about the signature here, but it's needed to conform to net/rpc requirements
func (s *Session) Register(args *SessionRegisterArgs, dummy *bool) error {
	if _, ok := sessionmap[args.SessionId]; ok {
		return fmt.Errorf("Session already exists: %s", args.SessionId)
	}
	sessionmap[args.SessionId] = &Session{args.Settings}
	return nil
}

// TODO(.): Not too happy about the signature here, but it's needed to conform to net/rpc requirements
func (s *Session) Unregister(SessionId string, dummy *bool) error {
	delete(sessionmap, SessionId)
	return nil
}
