package base

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/helper"
	"strings"
)

var db *gorm.DB
var repo domain.LoginUserRepo
var DbClose helper.Close

func init() {
	db, DbClose = helper.GetDb("authenticate")
	repo = &LoginUserRepoImpl{}
}

func GetRepo() domain.LoginUserRepo {
	return repo
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
	return "" + "_"
}

type LoginUserRepoImpl struct{}

func (l *LoginUserRepoImpl) GetOne(username string) domain.LoginUserDO {
	userDO, exist := l.FindOne(username)
	if !exist {
		panic("can not find")
	}
	return userDO
}

func (l *LoginUserRepoImpl) FindOne(username string) (userDO domain.LoginUserDO, exist bool) {
	db.Where("username=?", GetUsername(username)).First(&userDO)
	if userDO.ID == 0 {
		exist = false
	} else {
		exist = true
	}
	return userDO, exist
}

func (l *LoginUserRepoImpl) Add(userDO *domain.LoginUserDO) domain.LoginUserDO {
	userDO.Username = GetUsername(userDO.Username)
	result := db.Create(userDO)
	if result.Error != nil {
		panic(result.Error)
	} else {
		createUserDO := result.Value.(*domain.LoginUserDO)
		createUserDO.Username = InitUsername(createUserDO.Username)
		return *createUserDO
	}
}

func (l *LoginUserRepoImpl) Update() {
	panic("implement me")
}

func (l *LoginUserRepoImpl) FindSmsCode(mobile string) string {
	panic("implement me")
}
