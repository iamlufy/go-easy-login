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
	Insert(do *TenantDO) TenantDO
	FindByName(tenantName string) (tenant TenantDO, exist bool)
}
