package gosession

import "github.com/WindomZ/go-jwt/macro"

const (
	HEADER_KEY string = jwt.HEADER_KEY()          // Default: "Authorization"
	HEADER_VALUE_PREFIX = jwt.HEADER_VALUE_PREFIX() // Default: "Bearer"
)
