package middlewares

import (
    "time"
    "strings"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"  
    "github.com/InfamousFreak/Tech-Task-24/config"
)

func AuthMiddleware() fiber.Handler {
    jwtKey := []byte(config.Secret)
    return func(c *fiber.Ctx) error {
        // check if the JWT token is present in the Authorization header
        authHeader := c.Get("Authorization")
        var tokenString string
        if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
            tokenString = strings.TrimPrefix(authHeader, "Bearer ")
        } else {
            // if the token is not in the Authorization header, check for a cookie named "token"
            tokenString = c.Cookies("token")
            if tokenString == "" {
                return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                    "error":   true,
                    "message": "Missing or malformed JWT",
                })
            }
        }

        
        claims, err := parseAndVerifyToken(tokenString, jwtKey)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": true,
                "message": "Invalid or expired JWT: " + err.Error(),
            })
        }

        // token is valid, store the claims in the context
        c.Locals("claims", claims)
        return c.Next()
    }
}

func parseAndVerifyToken(tokenString string, jwtKey []byte) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token signing method")
        }
        return jwtKey, nil
    })

    if err != nil {
        return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token: "+err.Error())
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
    }

    if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
        return nil, fiber.NewError(fiber.StatusUnauthorized, "Token expired")
    }

    return claims, nil
}



