package kamar

import "github.com/gofiber/fiber/v2"

func RouterKamar(app *fiber.App) {
	router := app.Group("/api/1.0/kamar")
	router.Post("/", CreateKamar)
	router.Get("/", GetAllKamar)
	router.Get("/:id", GetKamarById)
	router.Get("/hotel/:id", GetKamarByIdHotel)
	router.Put("/:id", UpdateKamar)
	router.Delete("/:id", DeleteKamar)
}
