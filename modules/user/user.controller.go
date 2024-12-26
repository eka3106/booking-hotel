package user

import "github.com/gofiber/fiber/v2"

func RouterUser(app *fiber.App) {
	router := app.Group("/api/1.0")
	router.Post("/register", CreateUser)
	router.Post("/register-admin-hotel", CreateAdminHotel)
	router.Post("/login", Login)
	router.Post("/logout", Logout)
}
