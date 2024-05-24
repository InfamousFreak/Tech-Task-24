package models

import "gorm.io/gorm"

type Suggestion struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	MenuItemID uint `json:"menu_item_id"`
}
