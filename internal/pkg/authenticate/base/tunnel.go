package base

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
	"strings"
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

func GetUsername(username string) string {
	return getTenantCode() + username
}

func InitUsername(username string) string {
	if strings.HasPrefix(username, getTenantCode()) {
		return strings.Split(username, getTenantCode())[1]
	}
	return username
}

func getTenantCode() string {
	//TODO inject tenant code
	return "" + "_"
}

func (l PsqlTunnel) GetOne(username string) domain.LoginUserDO {
	userDO, exist := l.FindOne(username)
	if !exist {
		panic("can not find")
	}
	return userDO
}

func (l PsqlTunnel) FindOne(username string) (userDO domain.LoginUserDO, exist bool) {
	l.Where("username=?", GetUsername(username)).First(&userDO)
	if userDO.ID == 0 {
		exist = false
	} else {
		exist = true
	}
	return userDO, exist
}

func (l PsqlTunnel) Add(userDO *domain.LoginUserDO) domain.LoginUserDO {
	userDO.Username = GetUsername(userDO.Username)
	result := l.Create(userDO)
	if result.Error != nil {
		panic(result.Error)
	} else {
		createUserDO := result.Value.(*domain.LoginUserDO)
		createUserDO.Username = InitUsername(createUserDO.Username)
		return *createUserDO
	}
}

func (l PsqlTunnel) Update(model domain.LoginUserDO, updateFields map[string]interface{}) domain.LoginUserDO {
	return *l.Model(&model).Updates(updateFields).Value.(*domain.LoginUserDO)
}

func (l PsqlTunnel) FindSmsCode(mobile string) string {
	panic("implement me")
}
