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
)

// Returns the Session object associated with this Args or nil.
func (a *Args) Session() *Session {
	return sessionmap[a.SessionId]
}

// Returns the Session object associated with this Args or creates a new
// session id and associotes the Args object with that session if one does
// not exist
func (a *Args) SessionOrCreate(id string) *Session {
	if s := a.Session(); s != nil {
		return s
	} else {
		sessionmap[id] = &Session{*NewSettings()}
		a.SessionId = id
	}
	return a.Session()
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
