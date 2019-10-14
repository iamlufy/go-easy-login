package authenticate

import (
	"errors"
	"fmt"
	"oneday-infrastructure/internal/pkg/authenticate/base"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/internal/pkg/authenticate/facade"
	"oneday-infrastructure/tools"
)

var service = domain.InitLoginUserService(base.InitLoginUserRepo(tools.OpenDB))

func init() {
	if &service == nil {
		panic("LoginUserService should init ")
	}

}

func Login(cmd *domain.LoginCmd) (string, error) {
	//TODO check tenant
	userStatus := service.GetUserStatus(cmd.Username, cmd.TenantCode)
	if userStatus != domain.ALLOWED {
		return "", errors.New(fmt.Sprintf("userStatus is %s", userStatus))
	}
	matcherResult := service.Authenticate(cmd)
	if !matcherResult {
		return "", errors.New("match fail")
	}
	//TODO event
	return facade.GenerateToken(cmd.UniqueCode, cmd.EffectiveSeconds), nil
}

func AddUser(cmd *domain.AddLoginUserCmd) domain.AddUserResult {
	return service.AddUser(cmd)
}

func ReSetPassword(cmd *domain.ResetPasswordCmd) domain.ResetPasswordResult {
	return service.ReSetPassword(cmd)
}
