package domain

import "github.com/jinzhu/gorm"

type LoginUserDO struct {
	gorm.Model
	//TODO unique_index
	Username   string `gorm:"type:varchar(100);unique_index;not null"`
	Password   string `gorm:"type:varchar(100);unique_index;not null;default:''"`
	IsLock     bool   `gorm:"type:boolean;not null;default:false"`
	TenantCode string `gorm:"type:varchar(100);unique_index;not null"`
	Mobile     string `gorm:"type:varchar(100)"`
}
type LoginUserRepo interface {
	Add(*LoginUserDO) LoginUserDO
	GetOne(username string, tenantCode string) LoginUserDO
	Update(model LoginUserDO, updateFields map[string]interface{}) LoginUserDO
	FindOne(username string, tenantCode string) (userDO LoginUserDO, exist bool)
	FindSmsCode(mobile string) string
}

func updatePassword(repo LoginUserRepo, user LoginUserDO) LoginUserDO {
	return repo.Update(user, map[string]interface{}{"password": user.Password})
}

func findUser(repo LoginUserRepo, username, tenantCode string) (LoginUserDO, bool) {
	user, _ := repo.FindOne(username, tenantCode)
	return user, user.ID != 0
}
