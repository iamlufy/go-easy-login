package helper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"oneday-infrastructure/authenticate/domain"
	tenant_domain "oneday-infrastructure/tenant/domain"
	"time"
)

func GetDb(service string) *gorm.DB {
	//TODO get config by service
	//TODO connection pool or prevent to build connection every time
	db, err := gorm.Open(getConfig(""))
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&domain.LoginUserDO{}, &tenant_domain.TenantDO{})
	db.DB().SetConnMaxLifetime(time.Second / 2)
	return db
}

//TODO add to config.yml
func getConfig(service string) (string, string) {
	return "postgres", "host=localhost port=9000 user=oneday dbname=postgres password=123456 sslmode=disable"

}
