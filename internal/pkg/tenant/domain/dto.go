package domain

type AddTenantCmd struct {
	TenantName string
}
type TenantCO struct {
	TenantName string
	UniqueCode string
}

func ToTenantDO(cmd *AddTenantCmd) *TenantDO {
	return &TenantDO{
		TenantName: cmd.TenantName,
	}
}

func ToTenantCO(do TenantDO) TenantCO {
	return TenantCO{
		TenantName: do.TenantName,
		UniqueCode: do.UniqueCode,
	}
}

type AddTenantSuccess string

const AddSuccess AddTenantSuccess = "SUCCESS"

const TenantExist AddTenantSuccess = "EXISTED"
