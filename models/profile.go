package models

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model
	userID   uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
}
