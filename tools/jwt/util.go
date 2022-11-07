package jwt

import (
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"time"
)

func ParseNumericDate(t time.Time) *jwtV4.NumericDate {
	return jwtV4.NewNumericDate(t)
}
