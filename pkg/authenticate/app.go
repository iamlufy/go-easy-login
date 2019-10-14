package authenticate

import (
	"errors"
	"oneday-infrastructure/internal/pkg/authenticate/base"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/tools"
)

var service = func(tenantCode string) domain.LoginUserService {
	return domain.NewLoginUserService(base.NewLoginUserRepo(tools.OpenDB, tenantCode))
}

func init() {
	if &service == nil {
		panic("LoginUserService should init ")
	}
}

func Login(cmd *domain.LoginCmd, tenantCode string) (string, error) {
	//TODO check tenant

	token, matcherResult := service(tenantCode).Authenticate(cmd)
	if string(matcherResult) != domain.Success {
		return "", errors.New(string(matcherResult))
	}
	//TODO event
	return token, nil
}

func AddUser(cmd *domain.AddLoginUserCmd, tenantCode string) {
	repo := base.NewLoginUserRepo(tools.OpenDB, tenantCode)
	user := base.ToLoginUserDO(cmd)
	user.Password = service(tenantCode).Encrypt(cmd.EncryptWay, cmd.Password)
	repo.Add(&user)
}

func ReSetPassword(cmd *domain.ResetPasswordCmd, tenantCode string) domain.ResetPasswordResult {
	return service(tenantCode).ReSetPassword(cmd)
}
