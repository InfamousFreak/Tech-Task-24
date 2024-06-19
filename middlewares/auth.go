package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	//"github.com/dgrijalva/jwt-go"
    //"github.com/InfamousFreak/Tech-Task-24/config"
	//"strings"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler { //function only accepts an argument 'secret' which is a syring and returns a fiber.Handler which is a type alias for handling the http requests in the fiber framework
	return jwtware.New(jwtware.Config{ //return a new jwt middleware instance, calls jwtware.new with a config objet to create new middleware instance,
		SigningKey: []byte(secret), //sets the signingkey field of the config struct to the secret converted in the byte[] slice, is used bymiddlware to verify the signature of incoming jwt tokens
	})
}

/*func AuthMiddleware(secret string) fiber.Handler {
	jwtKey := []byte(secret)
	return func(c *fiber.Ctx) error {
		// First, check if the JWT token is present in the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse and verify the JWT token
			claims, err := parseAndVerifyToken(tokenString, jwtKey)
			if err != nil {
				return err
			}

			// Token is valid, store the claims in the context
			c.Locals("claims", claims)
			return c.Next()
		}

		// If the token is not in the Authorization header, check for a cookie named "token"
		accessToken := c.Cookies("token")
		if accessToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Missing or malformed JWT",
			})
		}

		// Parse and verify the JWT token
		claims, err := parseAndVerifyToken(accessToken, jwtKey)
		if err != nil {
			return err
		}

		// Token is valid, store the claims in the context
		c.Locals("claims", claims)
		return c.Next()
	}
}*/

/*func parseAndVerifyToken(tokenString string, jwtKey []byte) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   true,
					"message": "Malformed token",
				})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   true,
					"message": "Token expired",
				})
			} else {
				return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   true,
					"message": "Invalid token",
				})
			}
		} else {
			return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid token",
			})
		}
	}

	if !token.Valid {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid token",
		})
	}

	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Token expired",
		})
	}

	return claims, nil
}*/



