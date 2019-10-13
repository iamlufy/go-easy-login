package tenant

import (
	"github.com/jinzhu/gorm"
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

var tenantService = InitTenantService(
	base.InitTenantRepo(func(name string) *gorm.DB {
		return tools.OpenDB(name)
	}))

var genUniqueCode GenUniqueCode = func() string { return strconv.Itoa(rand.Intn(1000000)) }

func AddTenant(cmd *AddTenantCmd) (TenantCO, AddTenantSuccess) {
	return tenantService.AddTenant(cmd, genUniqueCode)

}
