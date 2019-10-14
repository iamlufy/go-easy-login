package tenant

import (
	"math/rand"
	"oneday-infrastructure/internal/pkg/tenant/base"
	. "oneday-infrastructure/internal/pkg/tenant/domain"
	"oneday-infrastructure/tools"
	"strconv"
)

func init() {
	if &tenantService == nil {
		panic("tenant service should init first")
	}
}

var repo = base.InitTenantRepo(tools.OpenDB)
var tenantService = InitTenantService(repo)

var genUniqueCode GenUniqueCode = func() string { return strconv.Itoa(rand.Intn(1000000)) }

func Add(cmd *AddTenantCmd) (TenantCO, AddTenantSuccess) {
	return tenantService.Add(cmd, genUniqueCode)

}
func IsExist(code string) bool {
	_, exist := repo.FindByCode(code)
	return exist
}
