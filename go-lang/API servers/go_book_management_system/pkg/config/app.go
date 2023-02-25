package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// db variable which is pointer to db
var (
	db *gorm.DB
)

func Connect() {

	dsn := "host=localhost user=awesome_dev password=password dbname=playground_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	pDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	println("connected to Db ... ")
	db = pDb
}

func GetDB() *gorm.DB {
	return db
}
