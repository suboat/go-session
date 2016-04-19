package gosession

import "net/http"

type SessionStore interface {
	Store(string, ...[]byte) error
	Path(string)
	New(*http.Request, string) (SessionStore, error)
	Get(*http.Request, string) (SessionStore, error)
	Save(*http.Request, http.ResponseWriter) error
	Name() string
	Flashes(...string) []interface{}
	AddFlash(interface{}, ...string)
}

const (
	SessionKeywordTag    string = "_tag_"
	SessionKeywordLevel         = "_level_"
	SessionKeywordStatus        = "_status_"
)
