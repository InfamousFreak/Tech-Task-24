package models

import "gorm.io/gorm"

type MenuItem struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Tags        string  `json:"tags"`
}
