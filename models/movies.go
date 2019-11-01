package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

type Movies struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Year  string `gorm:"column:year" json:"year"`
	Description   string `gorm:"column:description" json:"description"`
}

