package base

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	domain2 "oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/tools"
	"strings"
)

var repo Repo

type Repo struct {
	PsqlTunnel
}

type PsqlTunnel struct {
	*gorm.DB
}

func init() {
	repo = NewRepo()
}

func NewRepo() Repo {
	return Repo{
		PsqlTunnel{DB: tools.GetDb("authenticate")}}
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

func (l *PsqlTunnel) GetOne(username string) domain2.LoginUserDO {
	userDO, exist := l.FindOne(username)
	if !exist {
		panic("can not find")
	}
	return userDO
}

func (l *PsqlTunnel) FindOne(username string) (userDO domain2.LoginUserDO, exist bool) {
	l.Where("username=?", GetUsername(username)).First(&userDO)
	if userDO.ID == 0 {
		exist = false
	} else {
		exist = true
	}
	return userDO, exist
}

func (l *PsqlTunnel) Add(userDO *domain2.LoginUserDO) domain2.LoginUserDO {
	userDO.Username = GetUsername(userDO.Username)
	result := l.Create(userDO)
	if result.Error != nil {
		panic(result.Error)
	} else {
		createUserDO := result.Value.(*domain2.LoginUserDO)
		createUserDO.Username = InitUsername(createUserDO.Username)
		return *createUserDO
	}
}

func (l *PsqlTunnel) Update(model domain2.LoginUserDO, updateFields map[string]interface{}) domain2.LoginUserDO {
	return *l.Model(&model).Updates(updateFields).Value.(*domain2.LoginUserDO)
}

func (l *PsqlTunnel) FindSmsCode(mobile string) string {
	panic("implement me")
}
