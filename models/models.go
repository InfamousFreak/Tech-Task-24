package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
<<<<<<< HEAD
=======

type User struct {
	ID       int
	Email    string
	Password string
	City     string
}
>>>>>>> origin/master
