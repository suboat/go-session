package gosession

import (
	"github.com/WindomZ/go-jwt"
	"net/http"
)

type JWT struct {
}

func NewJWT(path string) (*JWT, error) {
	j := &JWT{}
	if err := j.Load(path); err != nil {
		return nil, err
	}
	return j, nil
}

func (j *JWT) Load(path string) error {
	return jwt.NewConfig(path).Effect()
}

func (j *JWT) Sign(kid string, arg interface{}, minutes int) (string, error) {
	return jwt.Sign(kid, arg, minutes)
}

func (j *JWT) Parse(token string) (interface{}, error) {
	return jwt.Parse(token)
}

func (j *JWT) SignRequest(r *http.Request, kid string, arg interface{}, minutes int) error {
	return jwt.SignRequest(r, kid, arg, minutes)
}

func (j *JWT) ParseRequest(r *http.Request) (interface{}, error) {
	return jwt.ParseRequest(r)
}
