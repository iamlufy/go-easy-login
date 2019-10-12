package authenticate

import (
	"errors"
	"fmt"
	domain2 "oneday-infrastructure/internal/pkg/authenticate/domain"
	facade2 "oneday-infrastructure/internal/pkg/authenticate/facade"
)

func Login(cmd *domain2.LoginCmd) (string, error) {

	userStatus := domain2.GetUserStatus(cmd.Username)
	if userStatus != domain2.ALLOWED {
		return "", errors.New(fmt.Sprintf("userStatus is %s", userStatus))
	}
	matcherResult := domain2.Authenticate(cmd)
	if !matcherResult {
		return "", errors.New("match fail")
	}
	code := domain2.GetUniqueCode(cmd.Username)
	//TODO event
	return facade2.GenerateToken(code, cmd.EffectiveSeconds), nil
}

func AddUser(cmd *domain2.AddLoginUserCmd) domain2.AddUserResult {
	return domain2.AddUser(cmd)
}

func ReSetPassword(cmd *domain2.ResetPasswordCmd) domain2.ResetPasswordResult {
	return domain2.ReSetPassword(cmd)
}
