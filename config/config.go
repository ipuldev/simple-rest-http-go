package config

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
