package token

import (
	"oneday-infrastructure/internal/pkg/token"
)

/**
1.check token if valid,it will refresh token if token is out of expired time
but user still active
2.it will return unique code ,which is set by @{Generate}
*/
func CheckToken(tokenString string) (string, bool) {
	return token.VerifyAndRefresh(tokenString)
}
