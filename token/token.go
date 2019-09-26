package token

import (
	"github.com/dgrijalva/jwt-go"
	"oneday-infrastructure/authenticate/base/cache"
	"time"
)

//TODO(zzf):to be configurable
const Second = 3600
const ActiveTime = 600

var tokenSecret = []byte("123")

func Generate(uniqueCode string, effectiveSeconds int) string {
	token := generateJwt(uniqueCode, int64(effectiveSeconds))
	cacheToken(uniqueCode, token)
	return token
}

func Verify(tokenString string) (bool, string) {
	token, e := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return tokenSecret, nil
	})
	if e != nil {
		claim := token.Claims.(jwt.MapClaims)
		code := claim["code"].(string)
		if getCache(code) != "" {
			return true, Generate(code, Second)
		} else {
			panic(e)
		}
	}

	return token.Valid, token.Raw

}

func generateJwt(uniqueCode string, effectiveSeconds int64) string {
	now := time.Now()
	//issue at time
	iat := now.Unix()
	// not before
	nbf := iat
	exp := now.Add(time.Duration(effectiveSeconds * 1e9)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"code": uniqueCode,
		"nbf":  nbf,
		"iat":  iat,
		"exp":  exp,
	})

	if tokenString, err := token.SignedString(tokenSecret); err == nil {
		return tokenString
	} else {
		panic(err)
	}
}

func cacheToken(key, token string) {
	cache.Add("token:"+key, token, ActiveTime)
}

func getCache(key string) string {
	return cache.Get("token:" + key)
}
