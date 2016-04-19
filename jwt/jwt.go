package gosession

import (
	"github.com/WindomZ/go-jwt"
	"net/http"
)

type JWT struct {
}

// New a JWT
// See JWT.Load
func NewJWT(path string) (*JWT, error) {
	j := &JWT{}
	if err := j.Load(path); err != nil {
		return nil, err
	}
	return j, nil
}

// Load key files
// path is a path of key files
// Can multi call
func (j *JWT) Load(path string) error {
	return jwt.NewConfig(path).Effect()
}

// Generate the signing string.
func (j *JWT) Sign(kid string, arg interface{}, minutes int) (string, error) {
	return jwt.Sign(kid, arg, minutes)
}

// Parse, validate, and return a token.
func (j *JWT) Parse(token string) (interface{}, error) {
	return jwt.Parse(token)
}

// Generate the signing string, and set into http request
func (j *JWT) SignRequest(r *http.Request, kid string, arg interface{}, minutes int) error {
	return jwt.SignRequest(r, kid, arg, minutes)
}

// Parse http request, validate, and return a token.
func (j *JWT) ParseRequest(r *http.Request) (interface{}, error) {
	return jwt.ParseRequest(r)
}
