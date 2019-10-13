package domain

type TenantService struct{}

func InitTenantService(tenantRepo TenantRepo) TenantService {
	InitTenantRepo(tenantRepo)
	return TenantService{}
}

type GenUniqueCode func() string

func (service TenantService) AddTenant(cmd *AddTenantCmd, genUniqueCode GenUniqueCode) (TenantCO, AddTenantSuccess) {
	if _, exist := find(cmd.TenantName); exist {
		return TenantCO{}, TenantExist
	}
	tenantDO := ToTenantDO(cmd)
	tenantDO.UniqueCode = genUniqueCode()
	return ToTenantCO(add(tenantDO)), AddSuccess
}
