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

func (psql TenantPsqlTunnel) Insert(do *domain.TenantDO) domain.TenantDO {
	result := psql.Create(do)
	if result.Error != nil {
		panic(result.Error)
	} else {
		return *result.Value.(*domain.TenantDO)
	}
}

func (psql TenantPsqlTunnel) FindByName(tenantName string) (tenant domain.TenantDO, exist bool) {
	psql.Where("tenant_name=?", tenantName).First(&tenant)
	return tenant, tenant.ID != 0

}

func (psql TenantPsqlTunnel) FindByCode(code string) (tenant domain.TenantDO, exist bool) {
	psql.Where("tenant_code = ?", code).First(&tenant)
	return tenant, tenant.ID != 0
}
