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

// TODO inject more elegant
func NewRepo(loginUserRepo LoginUserRepo) {
	repo = loginUserRepo
}

func Add(userDO *LoginUserDO) {
	repo.Add(userDO)
}

func FindUser(username string) (LoginUserDO, bool) {
	return repo.FindOne(username)
}

func GetUser(username string) LoginUserDO {
	return repo.GetOne(username)
}

func FindSmsCode(mobile string) string {
	return repo.FindSmsCode(mobile)

}
func GetUniqueCode(username string) string {
	return repo.GetOne(username).UniqueCode
}

func UpdatePassword(user LoginUserDO) LoginUserDO {
	return repo.Update(user, map[string]interface{}{"password": user.Password})
}
