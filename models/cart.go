package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model //is a struct provided by gorm to include common fields fordatabase mdoels

	UserID     uint `json:"user_id"` //type unsigned integer and how the field should look in the json format
	MenuItemID uint `json:"menu_item_id"`
	Quantity   int  `json:"quantity"`
}
