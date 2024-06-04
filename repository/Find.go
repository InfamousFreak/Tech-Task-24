package repository

import (
	"errors"

	"github.com/InfamousFreak/Tech-Task-24/models"
)

// simulate a database call
func Find(email, password string) (*models.UserProfile, error) { //a function that takes email and password as string as arguments, and returns a pointer to models.User and error which will be nil if no user is found
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
}
