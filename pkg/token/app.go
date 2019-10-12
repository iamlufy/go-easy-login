package token

import (
	token3 "oneday-infrastructure/internal/pkg/token"
)

func CheckLogin(token string) (string, bool) {
	return token3.VerifyAndRefresh(token)

}
