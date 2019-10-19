package main

import (
	"oneday-infrastructure/internal/pkg/tenant/base"
	"oneday-infrastructure/tools"
)

func main() {

	tools.OpenDB("tenant").
		AutoMigrate(base.TenantDO{})
}
