package domain

import "oneday-infrastructure/tools"

type TenantRepo interface {
	InsertTenant(tenant Tenant)
	FindByName(tenantName string) (tenant Tenant, exist bool)
	GetByCode(tenantCode string) (tenant Tenant)
	InsertUser(user User)
}

type TenantService struct {
	TenantRepo
}

func InitTenantService(tenantRepo TenantRepo) TenantService {
	return TenantService{tenantRepo}
}

type GenUniqueCode func() string

// TODO move to admin service
func (service TenantService) Add(cmd *AddTenantCmd, genUniqueCode GenUniqueCode) (TenantCO, AddTenantSuccess) {
	if tenant, exist := service.FindByName(cmd.TenantName); !exist {
		tenant.TenantCode = genUniqueCode()
		service.InsertTenant(tenant)
		return ToTenantCO(tenant), AddSuccess
	}
	return TenantCO{}, TenantNameExist

}

func (service TenantService) AddUser(cmd *AddUserCmd) {
	tenant := service.GetByCode(cmd.TenantCode)

	user := tenant.generateUser(
		cmd.Username,
		tools.ChooseEncrypter(cmd.EncryptWay)(cmd.Password),
		cmd.Mobile)

	service.InsertUser(user)

}
