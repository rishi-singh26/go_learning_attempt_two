package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "SomeSectedKey"

func GenerageToken(email string, userid int64) (string, error) {
	toekn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return toekn.SignedString([]byte(secretKey))
}
