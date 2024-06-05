package models

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model
	UserID   uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
	Preferences string `json:"preferences"`
}
