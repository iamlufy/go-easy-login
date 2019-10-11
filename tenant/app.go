package tenant

import "oneday-infrastructure/tenant/domain"

func Add(cmd *tenant_domain.AddTenantCmd) (tenant_domain.TenantCO, tenant_domain.AddTenantSuccess) {
	return tenant_domain.AddTenant(cmd)

}
