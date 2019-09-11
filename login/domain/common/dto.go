package common

import (
	"oneday-infrastructure/login/domain"
)

type AddLoginUserCmd struct {
	Username   string
	Password   string
	EncryptWay string
	TenantId   string
	UniqueCode string
	mobile     string
}
type LoginCmd struct {
	Username         string
	TenantId         string
	EffectiveSeconds int
	Mobile           string
	SourceCode       string
	LoginWay         string
	EncryptWay       string
}

func ToLoginUserDO(cmd *AddLoginUserCmd) *domain.LoginUserDO {
	return &domain.LoginUserDO{
		Username:   cmd.Username,
		Password:   cmd.Password,
		TenantId:   cmd.TenantId,
		UniqueCode: cmd.UniqueCode,
		IsLock:     true,
		Mobile:     cmd.mobile,
	}
}

func ToLoginUserE(dataObject domain.LoginUserDO) *domain.LoginUserE {

	return &domain.LoginUserE{
		Username:   dataObject.Username,
		IsLock:     dataObject.IsLock,
		UniqueCode: dataObject.UniqueCode,
		Mobile:     dataObject.Mobile,
	}

}
