package domain

import (
	"github.com/jinzhu/gorm"
)

type TenantDO struct {
	gorm.Model
	TenantName string `gorm:"type:varchar(20);unique_index;not null"`
	UniqueCode string `gorm:"type:varchar(50);unique_index;not null"`
}

type TenantRepo interface {
	Add(do *TenantDO) TenantDO
	FindByName(tenantName string) (t TenantDO)
}

var repo TenantRepo

func InitTenantRepo(tenantRepo TenantRepo) TenantRepo {
	repo = tenantRepo
	return repo
}

func getRepo() TenantRepo {
	// TODO inject more elegant
	return repo
}

func add(tenant *TenantDO) TenantDO {
	return getRepo().Add(tenant)
}

func find(tenantName string) (t TenantDO, exist bool) {
	tenantDO := getRepo().FindByName(tenantName)
	return tenantDO, tenantDO.ID != 0
}
