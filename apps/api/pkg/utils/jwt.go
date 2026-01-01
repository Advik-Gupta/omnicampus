package utils

import (
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func CreateJWT(userID string, email string) (string, error) {
	t := jwt.New()

	t.Set(jwt.SubjectKey, userID)
	t.Set("email", email)
	t.Set(jwt.IssuedAtKey, time.Now())
	t.Set(jwt.ExpirationKey, time.Now().Add(24*time.Hour))

	signed, err := jwt.Sign(t, jwt.WithKey(jwa.HS256, jwtKey))
	return string(signed), err
}

func JwtKey() []byte {
	return jwtKey
}