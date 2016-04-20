package gosession

import "github.com/WindomZ/go-jwt"

func IsTimeOutErr(err error) bool {
	return jwt.IsTimeOutErr(err)
}
