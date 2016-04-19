package gosession

import (
	. "github.com/suboat/go-session"
	"net/http"
)

func (s *Session) set(r *http.Request, w http.ResponseWriter, key string, value interface{}) error {
	if len(key) == 0 {
		return ErrKey
	} else if !s.hasSession(r) {
		return ErrSession
	}
	s.getSession().Values[key] = value
	return s.getSession().Save(r, w)
}

func (s *Session) get(r *http.Request, key string) interface{} {
	if len(key) == 0 {
	} else if !s.hasSession(r) {
	} else if v, ok := s.getSession().Values[key]; ok {
		return v
	}
	return nil
}

func (s *Session) SetTag(r *http.Request, w http.ResponseWriter, tag string) error {
	return s.set(r, w, SessionKeywordTag, tag)
}

func (s *Session) GetTag(r *http.Request) string {
	if v := s.get(r, SessionKeywordTag); v == nil {
	} else if tag, ok := v.(string); ok {
		return tag
	}
	return ""
}

func (s *Session) SetLevel(r *http.Request, w http.ResponseWriter, level uint32) error {
	return s.set(r, w, SessionKeywordLevel, level)
}

func (s *Session) GetLevel(r *http.Request) uint32 {
	if v := s.get(r, SessionKeywordLevel); v == nil {
	} else if level, ok := v.(uint32); ok {
		return level
	}
	return 0
}

func (s *Session) SetStatus(r *http.Request, w http.ResponseWriter, status uint64) error {
	return s.set(r, w, SessionKeywordStatus, status)
}

func (s *Session) GetStatus(r *http.Request) uint64 {
	if v := s.get(r, SessionKeywordStatus); v == nil {
	} else if status, ok := v.(uint64); ok {
		return status
	}
	return 0
}
