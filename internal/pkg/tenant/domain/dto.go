package domain

type AddTenantCmd struct {
	TenantName string `json:"tenantName" binding:"required,max=20"`
	TenantCode string `json:"tenantCode" `
}

type AddUserCmd struct {
	TenantCode string `json:"tenantCode" binding:"required"`
	Username   string `json:"username" binding:"required,max=20"`
	Mobile     string `json:"mobile" binding:"required,max=20"`
	Password   string `json:"password" binding:"required,max=20"`
	EncryptWay string `json:"encryptWay" binding:"required"`
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

func ToTenant(cmd *AddTenantCmd) Tenant {
	return Tenant{
		TenantName: cmd.TenantName,
		TenantCode: cmd.TenantCode,
	}
}

type AddTenantSuccess string

const AddSuccess AddTenantSuccess = "Success"

const TenantNameExist AddTenantSuccess = "TenantName_Existed"
