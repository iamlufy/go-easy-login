package base

import (
	"github.com/jinzhu/gorm"
	"oneday-infrastructure/internal/pkg/tenant/domain"
	"oneday-infrastructure/tools"
)

type TenantPsqlTunnel struct {
	*gorm.DB
}

type TenantRepo struct {
	TenantPsqlTunnel
}

var tenantRepo TenantRepo

func init() {
	tenantRepo = NewRepo()
}

func NewRepo() TenantRepo {
	return TenantRepo{
		TenantPsqlTunnel{DB: tools.GetDb("tenant")}}
}

func (psql TenantPsqlTunnel) Add(do *tenant_domain.TenantDO) tenant_domain.TenantDO {
	result := psql.Create(do)
	if result.Error != nil {
		panic(result.Error)
	} else {
		return *result.Value.(*tenant_domain.TenantDO)
	}
}

func (psql TenantPsqlTunnel) FindByName(tenantName string) (tenant tenant_domain.TenantDO) {
	psql.Where("tenant_name=?", tenantName).First(&tenant)
	return tenant

}
