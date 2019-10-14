package base

import "oneday-infrastructure/internal/pkg/authenticate/domain"

func ToLoginUser(do LoginUserDO) domain.LoginUser {
	return domain.LoginUser{
		Username: do.Username,
		Mobile:   do.Mobile,
		Password: do.Password,
		IsLock:   do.IsLock,
	}
}

func ToLoginUserDO(cmd *domain.AddLoginUserCmd) LoginUserDO {
	return LoginUserDO{
		Username: cmd.Username,
		Password: cmd.Password,
		IsLock:   false,
		Mobile:   cmd.Mobile,
	}
}
