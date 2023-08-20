package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error
var Db *gorm.DB

// connect database
func DbConnect() {
	dsn := "host=localhost user=takumi password=takumi dbname=todo port=5434 sslmode=disable TimeZone=Asia/Tokyo"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database !")
	}
}
