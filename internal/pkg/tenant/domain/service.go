package domain

type TenantService struct {
	TenantRepo
}

func InitTenantService(tenantRepo TenantRepo) TenantService {
	return TenantService{tenantRepo}
}

type GenUniqueCode func() string

func (service TenantService) Add(cmd *AddTenantCmd, genUniqueCode GenUniqueCode) (TenantCO, AddTenantSuccess) {
	if _, exist := service.FindByName(cmd.TenantName); exist {
		return TenantCO{}, TenantExist
	}
	tenantDO := ToTenantDO(cmd)
	tenantDO.UniqueCode = genUniqueCode()
	return ToTenantCO(service.Insert(tenantDO)), AddSuccess
}
