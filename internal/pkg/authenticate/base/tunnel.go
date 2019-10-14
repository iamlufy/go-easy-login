package base

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
)

type LoginUserRepo struct {
	PsqlTunnel
}

type PsqlTunnel struct {
	*gorm.DB
}

func InitLoginUserRepo(getDB func(string) *gorm.DB) LoginUserRepo {
	return LoginUserRepo{
		PsqlTunnel{DB: getDB("authenticate")}}
}

func (l PsqlTunnel) GetOne(username, tenantCode string) domain.LoginUserDO {
	userDO, _ := l.FindOne(username, tenantCode)
	if userDO.ID == 0 {
		panic("can not find")
	}
	return userDO
}

func (l PsqlTunnel) FindOne(username string, tenantCode string) (userDO domain.LoginUserDO, exist bool) {
	l.Where("username=? and tenant_code=?", username, tenantCode).First(&userDO)
	return userDO, userDO.ID != 0
}

func (l PsqlTunnel) Add(userDO *domain.LoginUserDO) domain.LoginUserDO {
	result := l.Create(userDO)
	if result.Error != nil {
		panic(result.Error)
	} else {
		createUserDO := result.Value.(*domain.LoginUserDO)
		return *createUserDO
	}
}

func (l PsqlTunnel) Update(model domain.LoginUserDO, updateFields map[string]interface{}) domain.LoginUserDO {
	return *l.Model(&model).Updates(updateFields).Value.(*domain.LoginUserDO)
}

func (l PsqlTunnel) FindSmsCode(mobile string) string {
	panic("implement me")
}
