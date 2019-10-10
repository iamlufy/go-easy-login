package authenticate

import (
	"errors"
	"fmt"
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/authenticate/facade"
)

func Login(cmd *domain.LoginCmd) (string, error) {

	userStatus := domain.GetUserStatus(cmd.Username)
	if userStatus != domain.ALLOWED {
		return "", errors.New(fmt.Sprintf("userStatus is %s", userStatus))
	}
	matcherResult := domain.Authenticate(cmd)
	if !matcherResult {
		return "", errors.New("match fail")
	}
	code := domain.GetUniqueCode(cmd.Username)
	//TODO event
	return facade.GenerateToken(code, cmd.EffectiveSeconds), nil
}

func AddUser(cmd *domain.AddLoginUserCmd) domain.AddUserResult {
	return domain.AddUser(cmd)
}

func SetPassword(cmd *domain.UpdatePasswordCmd) domain.UpdatePasswordResult {
	return domain.SetNewPassword(cmd)
}
