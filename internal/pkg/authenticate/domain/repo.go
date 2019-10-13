package domain

import "github.com/jinzhu/gorm"

type LoginUserDO struct {
	gorm.Model
	Username   string `gorm:"type:varchar(100);unique_index;not null"`
	Password   string `gorm:"type:varchar(100);unique_index;not null;default:''"`
	IsLock     bool   `gorm:"type:boolean;not null;default:false"`
	UniqueCode string `gorm:"type:varchar(100);unique_index;not null"`
	Mobile     string `gorm:"type:varchar(100)"`
}
type LoginUserRepo interface {
	Add(*LoginUserDO) LoginUserDO
	GetOne(username string) LoginUserDO
	Update(model LoginUserDO, updateFields map[string]interface{}) LoginUserDO
	FindOne(username string) (LoginUserDO, bool)
	FindSmsCode(mobile string) string
}

var repo LoginUserRepo

func InitLoginUserRepo(loginUserRepo LoginUserRepo) {
	repo = loginUserRepo
}

func getRepo() LoginUserRepo {
	return repo
}

func add(userDO *LoginUserDO) {
	getRepo().Add(userDO)
}

func findUser(username string) (LoginUserDO, bool) {
	return getRepo().FindOne(username)
}

func getUser(username string) LoginUserDO {
	return getRepo().GetOne(username)
}

func findSmsCode(mobile string) string {
	return getRepo().FindSmsCode(mobile)

}
func getUniqueCode(username string) string {
	return getRepo().GetOne(username).UniqueCode
}

func updatePassword(user LoginUserDO) LoginUserDO {
	return getRepo().Update(user, map[string]interface{}{"password": user.Password})
}
