package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"fmt"
	"brix-flix-backend/utils"
)

var db *gorm.DB //database
var err error

func Init() *gorm.DB {

	dbUri := utils.DbString()

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
		panic("Failed to connect to database")
	}

	db = conn
	conn.Debug().AutoMigrate(&Movie{}) //Database migration

	return conn
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

var DB = Init()