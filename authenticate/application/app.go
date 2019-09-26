package application

import (
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/token"
)

func Login(cmd *domain.LoginCmd) string {
	//TODO return detail reason
	if domain.CanLogin(cmd.Username) != domain.ALLOWED {
		return ""
	}
	matcherResult := domain.Authenticate(cmd)
	if !matcherResult {
		panic("match fail s")
	}
	code := domain.GetUniqueCode(cmd.Username)
	//TODO event
	return token.Generate(code, cmd.EffectiveSeconds)
}

func AddUser(cmd *domain.AddLoginUserCmd) {
	domain.AddUser(cmd)
}
