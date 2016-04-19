package gosession

import "net/http"

type SimpleSession interface {
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
