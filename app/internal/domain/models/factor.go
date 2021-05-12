package models

import (
	"gorm.io/gorm"
)

type Factor struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Star int `json:"star" form:"star"`
}
