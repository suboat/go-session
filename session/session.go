package gosession

import (
	"github.com/gorilla/sessions"
	. "github.com/suboat/go-session"
	"net/http"
	"time"
)

type Session struct {
	store   sessions.Store
	options *sessions.Options
	session *sessions.Session
}

func NewSession() *Session {
	return new(Session)
}

func newSession(s *Session, _s *sessions.Session) *Session {
	return &Session{store: s.store, session: _s}
}

func (s *Session) getStore() sessions.Store {
	if s.store == nil {
		s.Store("", []byte(time.Now().Format(time.RFC3339Nano)))
	}
	return s.store
}

func (s *Session) hasSession(r ...*http.Request) bool {
	if s.session == nil && len(r) != 0 {
		var err error
		s.session, err = s.getStore().Get(r[0], SessionKeyword)
		if err != nil {
			return false
		}
	}
	return (s.session != nil)
}

func (s *Session) getSession() *sessions.Session {
	return s.session
}

func (s *Session) Store(path string, keyPairs ...[]byte) error {
	if len(path) != 0 {
		store := sessions.NewFilesystemStore(path, keyPairs...)
		s.store = store
		s.options = store.Options
	} else {
		store := sessions.NewCookieStore(keyPairs...)
		s.store = store
		s.options = store.Options
	}
	return nil
}

func (s *Session) Path(path string) {
	s.options.Path = path
}

func (s *Session) New(r *http.Request, name string) (SessionStore, error) {
	_s, err := s.getStore().New(r, name)
	if err != nil {
		return nil, err
	}
	return newSession(s, _s), nil
}

func (s *Session) Get(r *http.Request, name string) (SessionStore, error) {
	_s, err := s.getStore().Get(r, name)
	if err != nil {
		return nil, err
	}
	return newSession(s, _s), nil
}

func (s *Session) Session(r *http.Request, name string) (err error) {
	s.session, err = s.getStore().Get(r, name)
	return nil
}

func (s *Session) Save(r *http.Request, w http.ResponseWriter) error {
	if s.hasSession(r) {
		return s.getStore().Save(r, w, s.getSession())
	}
	return ErrNil
}

func (s *Session) Name() string {
	if s.hasSession() {
		return s.getSession().Name()
	}
	return ""
}

func (s *Session) Flashes(vars ...string) []interface{} {
	if s.hasSession() {
		return s.getSession().Flashes(vars...)
	}
	return nil
}

func (s *Session) AddFlash(value interface{}, vars ...string) {
	if s.hasSession() {
		s.getSession().AddFlash(value, vars...)
	}
}
