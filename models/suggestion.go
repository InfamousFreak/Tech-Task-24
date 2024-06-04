package models

import "gorm.io/gorm"

type Suggestion struct {
	gorm.Model
	UserID     uint `json:"id"`
	MenuItemID uint `json:"item_id"`
}
