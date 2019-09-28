package token

func CheckLogin(token string) (string, bool) {
	return VerifyAndRefresh(token)

}
