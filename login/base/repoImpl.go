package base

import (
	"oneday-infrastructure/login/domain"
)

type LoginUserRepoImpl struct {
}

func (l *LoginUserRepoImpl) Add(*domain.LoginUserDO) {
	panic("implement me")
}

func (l *LoginUserRepoImpl) GetOne(username, tenantId string) *domain.LoginUserDO {
	panic("implement me")
}

func (l *LoginUserRepoImpl) Update() {
	panic("implement me")
}

func (l *LoginUserRepoImpl) FindSmsCode(mobile string) string {
	panic("implement me")
}
