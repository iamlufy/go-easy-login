package base

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oneday-infrastructure/authenticate/domain"
	"oneday-infrastructure/helper"
	"strings"
)

var repo domain.LoginUserRepo

func init() {
	NewRepo(func(name string) *gorm.DB {
		return helper.GetDb(name)
	})
}

func NewRepo(getDb func(name string) *gorm.DB) domain.LoginUserRepo {
	db := getDb("authenticate")
	repo = &LoginUserRepoImpl{DB: db}
	return repo
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
	//TODO inject tenant code
	return "" + "_"
}

type LoginUserRepoImpl struct {
	*gorm.DB
}

func (l *LoginUserRepoImpl) FindByUsernameAndPassword(username, password string) (userDO domain.LoginUserDO, exist bool) {
	l.Where("username=? and password=?", GetUsername(username), password).First(&userDO)
	if userDO.ID == 0 {
		exist = false
	} else {
		exist = true
	}
	return userDO, exist
}

func (l *LoginUserRepoImpl) GetOne(username string) domain.LoginUserDO {
	userDO, exist := l.FindOne(username)
	if !exist {
		panic("can not find")
	}
	return userDO
}

func (l *LoginUserRepoImpl) FindOne(username string) (userDO domain.LoginUserDO, exist bool) {
	l.Where("username=?", GetUsername(username)).First(&userDO)
	if userDO.ID == 0 {
		exist = false
	} else {
		exist = true
	}
	return userDO, exist
}

func (l *LoginUserRepoImpl) Add(userDO *domain.LoginUserDO) domain.LoginUserDO {
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

func (l *LoginUserRepoImpl) Update() {
	panic("implement me")
}

func (l *LoginUserRepoImpl) FindSmsCode(mobile string) string {
	panic("implement me")
}
