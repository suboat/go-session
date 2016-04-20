package gosession

import "github.com/WindomZ/go-jwt/macro"

// Default: "Authorization"
func HEADER_KEY() string {
	return jwt.HEADER_KEY()
}

// Default: "Bearer"
func HEADER_VALUE_PREFIX() string {
	return jwt.HEADER_VALUE_PREFIX()
}
