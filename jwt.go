package gosession

import "net/http"

type JWTToken interface {
	Load(string) error
	Sign(string, interface{}, int) (string, error)
	Parse(string) (interface{}, error)
	SetHTTPHeaderKey(key string)
	SignRequest(*http.Request, string, interface{}, int) error
	ParseRequest(*http.Request) (interface{}, error)
}
