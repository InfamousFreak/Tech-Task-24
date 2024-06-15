package repository

import (
	"errors"
	"gorm.io/gorm"
	"github.com/InfamousFreak/Tech-Task-24/models"
	"golang.org/x/crypto/bcrypt"

)

// simulate a database call
/*func Find(email, password string) (*models.UserProfile, error) { //a function that takes email and password as string as arguments, and returns a pointer to models.User and error which will be nil if no user is found
	//here you would query your database for the user with the given email
	if email == "test@gmail.com" && password == "test12345" {
		return &models.UserProfile{
			Name:     "Smarak",
			Email:    "test@gmail.com",
			Password: "test12345",
			City:     "Bhubaneshwar",
		}, nil
	}
	return nil, errors.New("user not found")
}*/


func Find(db *gorm.DB, email, password string) (*models.UserProfile, error) {
    var user models.UserProfile

    // Query the database for the user with the given email
    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    // Compare the provided password with the stored hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid password")
    }

    return &user, nil
}

