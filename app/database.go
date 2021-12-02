package app

import (
	"github.com/andikanugraha11/go-boilerplate/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=go_boiler_plate port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(entity.Order{})

	return db
}
