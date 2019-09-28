package application

import (
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/authenticate/facade"
)

func Login(cmd *domain.LoginCmd) string {

	if domain.CanLogin(cmd.Username) != domain.ALLOWED {
		return ""
	}
	matcherResult := domain.Authenticate(cmd)
	if !matcherResult {
		panic("match fail s")
	}
	code := domain.GetUniqueCode(cmd.Username)
	//TODO event
	return facade.GenerateToken(code, cmd.EffectiveSeconds)
}

func AddUser(cmd *domain.AddLoginUserCmd) {

	domain.AddUser(cmd, func(username string) bool {
		_, exist := domain.FindUser(username)
		return exist
	})
}
