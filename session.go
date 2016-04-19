package gosession

import "net/http"

type SessionStore interface {
	Store(string, ...[]byte) error
	Path(string)
	New(*http.Request, string) (SessionStore, error)
	Get(*http.Request, string) (SessionStore, error)
	Session(*http.Request, string) error
	Save(*http.Request, http.ResponseWriter) error
	Name() string
	Flashes(...string) []interface{}
	AddFlash(interface{}, ...string)
}

const SessionKeyword string = "_session"

type SessionCommonStore interface {
	SessionStore
	SetTag(*http.Request, http.ResponseWriter, string) error
	GetTag(*http.Request) string
	SetLevel(*http.Request, http.ResponseWriter, uint32) error
	GetLevel(*http.Request) uint32
	SetStatus(*http.Request, http.ResponseWriter, uint64) error
	GetStatus(*http.Request) uint64
}

const (
	SessionKeywordTag    string = "_tag"
	SessionKeywordLevel         = "_level"
	SessionKeywordStatus        = "_status"
)
