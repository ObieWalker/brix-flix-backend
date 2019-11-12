package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"fmt"
	"brix-flix-backend/utils"
)

var db *gorm.DB //database

func Init() {

	dbUri := utils.DbString()

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
		panic("Failed to connect to database")
	}

	db = conn
	db.Debug().AutoMigrate(&Movie{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
