package tenant_domain

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}
func AddTenant(cmd *AddTenantCmd) (TenantCO, AddTenantSuccess) {
	if _, exist := find(cmd.TenantName); exist {
		return TenantCO{}, TenantExist
	}
	tenantDO := ToTenantDO(cmd)
	// TODO make more sense
	tenantDO.UniqueCode = strconv.Itoa(rand.Intn(1000000))
	return ToTenantCO(add(tenantDO)), AddSuccess
}
