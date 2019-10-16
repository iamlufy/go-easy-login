package domain

type AddTenantCmd struct {
	TenantName string
}

type AddUserCmd struct {
	TenantCode string
	Username   string
	Mobile     string
	Password   string
	EncryptWay string
}

type TenantCO struct {
	TenantName string
	TenantCode string
}

func ToTenantCO(t Tenant) TenantCO {
	return TenantCO{
		TenantName: t.TenantName,
		TenantCode: t.TenantCode,
	}
}

type AddTenantSuccess string

const AddSuccess AddTenantSuccess = "Success"

const TenantNameExist AddTenantSuccess = "TenantName_Existed"
