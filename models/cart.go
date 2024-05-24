package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	MenuItemID uint `json:"menu_item_id"`
	Quantity   int  `json:"quantity"`
}
