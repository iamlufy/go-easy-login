package base

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oneday-infrastructure/authenticate/domain"
)

var repo *gorm.DB

func init() {
	//TODO
	db, err := gorm.Open("postgres", "host=localhost port=9000 user=oneday dbname=easy-login password=123456")
	repo = db
	if err != nil {
		panic("failed to connect database")
	}
}

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
