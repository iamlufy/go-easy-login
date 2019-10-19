package base

import "oneday-infrastructure/internal/pkg/tenant/domain"

func ToTenant(do *TenantDO) domain.Tenant {
	return domain.Tenant{
		TenantCode: do.TenantCode,
		TenantName: do.TenantName,
	}
}

func ToTenantDO(tenant *domain.Tenant) *TenantDO {
	return &TenantDO{
		TenantCode: tenant.TenantCode,
		TenantName: tenant.TenantName,
	}
}
