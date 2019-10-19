package base

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/tools"
)

type LoginUserDO struct {
	gorm.Model
	//TODO unique_index
	Username   string `gorm:"type:varchar(100);unique_index;not null"`
	Password   string `gorm:"type:varchar(100);unique_index;not null;default:''"`
	IsLock     bool   `gorm:"type:boolean;not null;default:false"`
	TenantCode string `gorm:"type:varchar(100);not null"`
	Mobile     string `gorm:"type:varchar(100)"`
}

func (LoginUserDO) TableName() string {
	return "loginUsers"
}

type LoginUserRepo struct {
	TenantCode string
	PsqlTunnel
}

type PsqlTunnel struct {
	*gorm.DB
}

var db *gorm.DB

func NewLoginUserRepo(getDB func(string) *gorm.DB, tenantCode string) LoginUserRepo {
	if db == nil {
		//TODO 控制并发
		db = getDB("authenticate")
	}
	return LoginUserRepo{
		tenantCode,
		PsqlTunnel{DB: db}}
}

func (psql PsqlTunnel) tenantQuery(tenantCode string) *gorm.DB {
	return psql.Where("tenant_code = ?", tenantCode)
}

func (psql PsqlTunnel) GetOne(username, tenantCode string) LoginUserDO {
	userDO, exist := psql.FindOne(username, tenantCode)
	if !exist {
		panic("can not find")
	}
	return userDO
}

func (psql PsqlTunnel) FindOne(username string, tenantCode string) (userDO LoginUserDO, exist bool) {
	psql.tenantQuery(tenantCode).Where("username =?", username).First(&userDO)
	return userDO, userDO.ID != 0
}

func (psql PsqlTunnel) Insert(userDO *LoginUserDO) {
	result := psql.Create(&userDO)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (psql PsqlTunnel) Update(model *LoginUserDO, updateFields map[string]interface{}) {
	psql.Model(&model).Updates(updateFields)
}

func (psql PsqlTunnel) UpdateFields(username string, tenantCode string, updateFields map[string]interface{}) LoginUserDO {
	userDO := psql.GetOne(username, tenantCode)
	return *psql.tenantQuery(tenantCode).Model(&userDO).Updates(updateFields).Value.(*LoginUserDO)
}

func (repo LoginUserRepo) FindSmsCode(mobile string) string {
	panic("implement me")
}
func (repo LoginUserRepo) GetOne(username string) domain.LoginUser {
	return ToLoginUser(repo.PsqlTunnel.GetOne(username, repo.TenantCode))
}

func (repo LoginUserRepo) UpdatePasswordByUsername(user domain.LoginUser) domain.LoginUser {
	return ToLoginUser(
		repo.PsqlTunnel.UpdateFields(
			user.Username,
			repo.TenantCode,
			tools.NewMap("password", user.Password)))
}

func (repo LoginUserRepo) FindOne(username string) (domain.LoginUser, bool) {
	userDO, exist := repo.PsqlTunnel.FindOne(username, repo.TenantCode)
	return ToLoginUser(userDO), exist
}
