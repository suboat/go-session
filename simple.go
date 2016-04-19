package gosession

import "net/http"

type SimpleSession interface {
	SessionStore
	SetTag(*http.Request, string) error
	GetTag(*http.Request) (string, error)
	SetLevel(*http.Request, uint32) error
	GetLevel(*http.Request) (uint32, error)
	SetStatus(*http.Request, uint64) error
	GetStatus(*http.Request) (uint64, error)
}
