package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Load(key string) string {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	//return the value of the variable
	return os.Getenv(key)
}

// secret key used to sign the JWT, this must be a secure key and should not be stored in the code
const Secret = "secret"
const AdminSecret = "secret"

type GConfig struct {
	GoogleLoginConfig oauth2.Config //defines a struct named config with a field googleloginconfig of type oauth2.config
	//GitHubLoginConfig oauth2.Config
}

var AppConfig GConfig //global variable of type config

func GoogleConfig() oauth2.Config { //function googleconfig returns a struct oauth2.config
	err := godotenv.Load(".env") //loads up the .env file
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err) //if there is some error in loading the env file then it stops executing the program
	}

	AppConfig.GoogleLoginConfig = oauth2.Config{ //field and struct
		RedirectURL:  "http://localhost:8080/google_callback", //redirect the user to the given url
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", //bunch of scope requests from the users the email and profile of the user
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint, //google's default endpoint
	}

	return AppConfig.GoogleLoginConfig //returns the struct
}
