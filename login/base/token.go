package base

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


func generateJwt(uniqueCode string, effectiveSeconds int) string {
	iat := time.Now().Unix()
	nbf := iat
	exp := time.Now().Add(time.Duration(time.Second.Seconds() * float64(effectiveSeconds))).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"code": uniqueCode,
		"nbf":  nbf,
		"iat":  iat,
		"exp":  exp,
	})
	//todo secret should to be config
	if tokenString, err := token.SignedString([]byte("123")); err == nil {
		return tokenString
	} else {
		panic(err)
	}
}

