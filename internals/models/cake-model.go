package models

import (
	"gorm.io/gorm"
)

type Cake struct {
	gorm.Model
	Size  string  `json:"size"`
	Price float64 `json:"price"`
}
