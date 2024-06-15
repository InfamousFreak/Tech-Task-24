package models

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model
	UserID   uint   `gorm:"unique" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
	Preferences string `json:"preferences"`
	Role string `json:"role"`
}

type Admin struct {
    gorm.Model
    UserProfileID       uint   `gorm:"primaryKey" json:"user_profile_id"`
    UserProfile         UserProfile `gorm:"foreignKey:UserProfileID;references:UserID"`
    BusinessLicenseNumber string `json:"business_license_number"`
    PhoneNumber         string `json:"phone_number"`
    RestaurantType      string `json:"restaurant_type"`
}
