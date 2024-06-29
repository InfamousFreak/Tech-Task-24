package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)

type LoginRequest struct {
    Email           string `json:"email"`
    Password        string `json:"password"`
    BusinessLicense string `json:"business_license,omitempty"`
}

type LoginResponse struct {
    Token string `json:"token"`
    UserID uint `json:"userId"`
}


type UserProfile struct {
    gorm.Model
    Name           string `json:"name"`
    Email          string `gorm:"unique;not null" json:"email"` // Ensure email is unique and not null
    Password       string `json:"password"`
    City           string `json:"city"`
    Role           string `json:"role"`
    BusinessLicense string `json:"business_license"`
}



type Admin struct {
    gorm.Model
    UserProfileID       uint   `gorm:"primaryKey" json:"user_profile_id"`
    UserProfile         UserProfile `gorm:"foreignKey:UserProfileID;references:UserID"`
    BusinessLicenseNumber string `json:"business_license_number"`
    RestaurantType      string `json:"restaurant_type"`
}

type Roles struct {
    Role            string `json:"role"`
    BusinessLicense string `json:"business_license,omitempty"`
}

type Claims struct {
	jwt.StandardClaims
}

type MenuItem struct {
    gorm.Model
	MenuItemID   uint   `gorm:"primaryKey" json:"item_id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Tags        string  `json:"tags"`
    ImageUrl    string  `json:"imageUrl"`
}

type CartItem struct {
    gorm.Model
    UserID     uint `gorm:"primaryKey" json:"user_id"`
    MenuItemID uint `gorm:"primaryKey" json:"item_id"`
    Quantity   int  `json:"quantity"`
}