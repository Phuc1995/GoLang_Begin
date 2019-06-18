package handler

type SessionStore interface {
	Find(string) (*Session, error)
	Save(*Session) error
	Delete(*Session) error
}

type FileSessionStore struct {
	filename string
	Session map[string]Session
}

/*func NewFileSessionStore(name string) (*FileSessionStore, error)  {
	store := &FileSessionStore{
		Session: map[string]Session{},
		filename: name,
	}
}*/