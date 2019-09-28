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
	return token
}

func Verify(tokenString string) (*jwt.Token, error) {
	//log error
	token, e := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return tokenSecret, nil
	})
	if token != nil {
		return token, nil
	} else {
		return nil, e
	}

}

func VerifyAndRefresh(tokenString string) (string, bool) {
	token, e := Verify(tokenString)
	if e != nil {
		//TODO fake token cause error,log
		return "", false
	}
	claim := token.Claims.(jwt.MapClaims)
	code := claim["code"].(string)
	if token.Valid {
		cacheToken(code, tokenString)
		return tokenString, true
	} else {
		if getCache(code) != "" {
			cacheToken(code, tokenString)
			return Generate(code, Second), true
		} else {
			return "", false
		}
	}

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
