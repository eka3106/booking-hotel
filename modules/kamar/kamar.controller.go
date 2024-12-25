package kamar

import "github.com/gofiber/fiber/v2"

func RouterKamar(app *fiber.App) {
	router := app.Group("/api/v1/kamar")
	router.Post("/", CreateKamar)
	router.Get("/", GetAllKamar)
	router.Get("/:id", GetKamarById)
	router.Put("/:id", UpdateKamar)
	router.Delete("/:id", DeleteKamar)
}
