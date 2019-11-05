package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"fmt"
	"brix-flix-backend/utils"
)

var db *gorm.DB //database
var err error

func Init() {

	dbUri := utils.DbString()
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err.Error())
		panic("Failed to connect to database")
	}
	defer conn.Close()

	conn.Debug().AutoMigrate(&Movie{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}