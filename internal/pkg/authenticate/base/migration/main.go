package main

import (
	"oneday-infrastructure/internal/pkg/authenticate/base"
	"oneday-infrastructure/tools"
)

func main() {

	tools.OpenDB("authenticate").AutoMigrate(base.LoginUserDO{}).
		AddIndex("uiq_tenant_code_username", "tenant_code", "username")
}
