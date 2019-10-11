package authenticate

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"oneday-infrastructure/authenticate/base"
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/authenticate/facade"
	"oneday-infrastructure/helper"
)

func init() {
	domain.NewRepo(base.NewRepo(func(name string) *gorm.DB {
		return helper.GetDb(name)
	}))
}

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

func ReSetPassword(cmd *domain.ResetPasswordCmd) domain.ResetPasswordResult {
	return domain.ReSetPassword(cmd)
}
