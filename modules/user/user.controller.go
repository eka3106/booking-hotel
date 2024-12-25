package user

import "github.com/gofiber/fiber/v2"

func RouterUser(app *fiber.App) {
	router := app.Group("/auth/v1")
	router.Post("/register", CreateUser)
	router.Post("/login", Login)
}
