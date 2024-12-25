package hotel

import "github.com/gofiber/fiber/v2"

func RouterHotel(app *fiber.App) {

	router := app.Group("/api/1.0/hotel")
	router.Post("/", CreateHotel)
	router.Get("/", GetAllHotel)
	router.Get("/:id", GetHotelById)
	router.Put("/:id", UpdateHotel)
	router.Delete("/:id", DeleteHotel)
}
