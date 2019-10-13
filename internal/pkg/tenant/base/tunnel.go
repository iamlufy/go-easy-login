package base

import (
	"github.com/jinzhu/gorm"
	"oneday-infrastructure/internal/pkg/tenant/domain"
)

type TenantPsqlTunnel struct {
	*gorm.DB
}

type TenantRepo struct {
	TenantPsqlTunnel
}

func InitTenantRepo(getDB func(string) *gorm.DB) TenantRepo {
	return TenantRepo{
		TenantPsqlTunnel{DB: getDB("tenant")}}
}

func (psql TenantPsqlTunnel) Add(do *domain.TenantDO) domain.TenantDO {
	result := psql.Create(do)
	if result.Error != nil {
		panic(result.Error)
	} else {
		return *result.Value.(*domain.TenantDO)
	}
}

func (psql TenantPsqlTunnel) FindByName(tenantName string) (tenant domain.TenantDO) {
	psql.Where("tenant_name=?", tenantName).First(&tenant)
	return tenant

}
