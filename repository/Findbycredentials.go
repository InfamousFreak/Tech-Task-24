package repository

import (
	"errors"

	"github.com/InfamousFreak/Tech-Task-24/models"
)

// simulate a database call
func Findbycredentials(email, password string) (*models.User, error) {
	//here you would query your database for the user with the given email
	if email == "test@gmail.com" && password == "test12345" {
		return &models.User{
			ID:       1,
			Email:    "test@gmail.com",
			Password: "test12345",
			City:     "Bhubaneshwar",
		}, nil
	}
	return nil, errors.New("user not found")
}
