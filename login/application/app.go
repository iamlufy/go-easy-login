package application

import (
	"oneday-infrastructure/login/base"
	"oneday-infrastructure/login/domain/common"
	"oneday-infrastructure/login/domain/service"
)

var loginService *service.LoginService

func init() {
	loginService = service.NewLoginService(&base.LoginUserRepoImpl{}, &base.TokenServiceImpl{})
}

func Login(cmd *common.LoginCmd) string {
	token, err := loginService.Login(cmd)

	if err != nil {
		panic(err)
	}
	//TODO event
	return token
}

func AddUser(cmd *common.AddLoginUserCmd) {
	loginService.AddUser(cmd)
}

func CheckLogin(token string) (bool, string) {
	return loginService.CheckLogin(token)
}
