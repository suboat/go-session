package gosession

import "net/http"

type JWTToken interface {
	Load(string) error
	Sign(string, interface{}, int)
	Parse(tokenString string) (interface{}, error)
	SignRequest(*http.Request, string, interface{}, int) error
	ParseRequest(*http.Request) (interface{}, error)
}
