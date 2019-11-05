package models

import (
	"brix-flix-backend/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Movie struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Year  string `gorm:"column:year" json:"year"`
	Description   string `gorm:"column:description" json:"description"`
}

func (movie *Movie) Create() (map[string] interface{}) {
	GetDB().Create(movie)

	if movie.ID <= 0 {
		return utils.Message(false, "Failed to create movie, connection error.")
	}

	response := utils.Message(true, "Movie has been added.")
	response["movie"] = movie
	return response
}