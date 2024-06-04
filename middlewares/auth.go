package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler { //function only accepts an argument 'secret' which is a syring and returns a fiber.Handler which is a type alias for handling the http requests in the fiber framework
	return jwtware.New(jwtware.Config{ //return a new jwt middleware instance, calls jwtware.new with a config objet to create new middleware instance,
		SigningKey: []byte(secret), //sets the signingkey field of the config struct to the secret converted in the byte[] slice, is used bymiddlware to verify the signature of incoming jwt tokens
	})
}
