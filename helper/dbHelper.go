package helper

import (
	"github.com/jinzhu/gorm"
	"oneday-infrastructure/authenticate/domain"
	"time"
)

func GetDb(service string) *gorm.DB {
	//TODO get config by service
	//TODO connection pool or prevent to build connection every time
	db, err := gorm.Open(getConfig(""))
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&domain.LoginUserDO{})
	db.DB().SetConnMaxLifetime(time.Second / 2)
	return db
}

//TODO add to config.yml
func getConfig(service string) (string, string) {
	return "postgres", "host=localhost port=9000 user=oneday dbname=postgres password=123456 sslmode=disable"

}
