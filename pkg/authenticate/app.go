package authenticate

import (
	"errors"
	"oneday-infrastructure/internal/pkg/authenticate/base"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/tools"
)

var service = func(tenantCode string) domain.LoginUserService {
	//TODO check tenantCode
	return domain.NewLoginUserService(base.NewLoginUserRepo(tools.OpenDB, tenantCode))
}

func init() {
	if &service == nil {
		panic("LoginUserService should init ")
	}
}

func Login(cmd *domain.LoginCmd, tenantCode string) (string, error) {
	//TODO check tenant

	token, authenticateResult := service(tenantCode).Authenticate(cmd)
	if authenticateResult.IsSuccess() {
		return "", errors.New(string(authenticateResult))
	}
	//TODO event
	return token, nil
}

func ReSetPassword(cmd *domain.ResetPasswordCmd, tenantCode string) domain.ResetPasswordResult {
	return service(tenantCode).ReSetPassword(cmd)
}
