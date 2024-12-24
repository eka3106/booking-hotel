package middleware

import (
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func AuthJWT() fiber.Handler {
	return verifyToken
}

func verifyToken(c *fiber.Ctx) error {
	bearer := c.Get("Authorization")
	if bearer == "" {
		return c.Next()

	}
	token := libs.ExtractToken(c)
	claims := &user.Claims{}
	// verify token
	secret := []byte(viper.GetString("SECRET_JWT"))
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return c.Next()

	}
	isExist := user.CheckToken(token, claims.IdUser)
	if !isExist {
		return c.Next()

	}
	c.Locals("user", claims)
	return c.Next()
}
