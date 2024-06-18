package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserProfile struct {
	gorm.Model
	UserID   uint   `gorm:"unique" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
	Role 	 string `json:"role"`
}

type Admin struct {
    gorm.Model
    UserProfileID       uint   `gorm:"primaryKey" json:"user_profile_id"`
    UserProfile         UserProfile `gorm:"foreignKey:UserProfileID;references:UserID"`
    BusinessLicenseNumber string `json:"business_license_number"`
    RestaurantType      string `json:"restaurant_type"`
}

type Role struct{
	Role string `json:"role"`
}

type Claims struct {
	jwt.StandardClaims
}

type MenuItem struct {
    gorm.Model
	MenuItemID   uint   `gorm:"primaryKey" json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Tags        string  `json:"tags"`
}

type CartItem struct {
	gorm.Model //is a struct provided by gorm to include common fields fordatabase mdoels

	UserID     uint `gorm:"primaryKey" json:"id"` //type unsigned integer and how the field should look in the json format
	MenuItemID uint `gorm:"primaryKey" json:"item_id"`
	Quantity   int  `json:"quantity"`
}