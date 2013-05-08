package content

type (
	Args struct {
		SessionId        string
		SessionOverrides Settings
	}

	CompleteAtArgs struct {
		Args
		Location SourceLocation
	}

	CompleteArgs struct {
		Args
		Location FullyQualifiedName
	}

	Session struct {
		Settings Settings
	}
)

// Returns the Session object associated with this Args or nil.
func (a *Args) Session() *Session {
	if a.SessionId != "" {
		// TODO actually look up session and return it
	}
	return nil
}

// If a Session is associated with this Args, then the Session's
// Settings is cloned and the Args's Settings are merged into the
// cloned Settings object. This to allow Argss to be "slim", i.e.
// only contain the keys of the specific settings it wants to override.
//
// If there's no Session nor Args Settings, the empty Settings object
// is returned.
func (it *Args) Settings() *Settings {
	if session := it.Session(); session != nil {
		set := session.Settings.Clone()
		set.Merge(&it.SessionOverrides)
		return set
	} else {
		return &it.SessionOverrides
	}
}
