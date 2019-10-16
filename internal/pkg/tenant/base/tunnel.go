package base

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"oneday-infrastructure/internal/pkg/authenticate/base"
	"oneday-infrastructure/internal/pkg/tenant/domain"
	"oneday-infrastructure/tools"
)

type PsqlTunnel struct {
	*gorm.DB
}

type TenantRepo struct {
	PsqlTunnel
}

type TenantDO struct {
	gorm.Model
	TenantName string `gorm:"type:varchar(20);unique_index;not null"`
	TenantCode string `gorm:"type:varchar(50);unique_index;not null"`
}

func NewTenantRepo(getDB func(string) *gorm.DB) TenantRepo {
	return TenantRepo{
		PsqlTunnel{DB: getDB("tenant")}}
}

func (psql PsqlTunnel) Insert(do *TenantDO) {
	result := psql.Create(&do)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (psql PsqlTunnel) FindOne(where map[string]interface{}) (tenant TenantDO, exist bool) {
	psql.Where(where).First(&tenant)
	exist = tenant.ID != 0
	return
}

func (psql PsqlTunnel) GetOne(where map[string]interface{}) (tenant *TenantDO) {
	psql.Where(where).First(&tenant)
	if tenant == nil {
		panic(fmt.Errorf("can not get user by %s ", where))
	}
	return tenant
}

func (repo TenantRepo) FindByName(tenantName string) (domain.Tenant, bool) {
	tenantDO, existed := repo.FindOne(map[string]interface{}{"tenant_name": tenantName})
	return ToTenant(tenantDO), existed
}

func (repo TenantRepo) GetByCode(code string) domain.Tenant {
	return ToTenant(*repo.GetOne(map[string]interface{}{"tenant_code": code}))
}

func (repo TenantRepo) InsertTenant(tenant domain.Tenant) {
	tenantDO := ToTenantDO(tenant)
	repo.PsqlTunnel.Insert(tenantDO)
}

func (repo TenantRepo) InsertUser(user domain.User) {
	base.NewLoginUserRepo(tools.OpenDB, user.TenantCode).Insert(
		&base.LoginUserDO{
			Username:   user.Username,
			Password:   user.Password,
			IsLock:     false,
			TenantCode: user.TenantCode,
			Mobile:     user.Mobile,
		})
}
