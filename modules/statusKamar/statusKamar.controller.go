package statuskamar

import "github.com/gofiber/fiber/v2"

func RouterStatusKamar(app *fiber.App) {
	router := app.Group("/api/1.0/status-kamar")
	router.Get("/", GetAllStatusKamar)
	router.Post("/", CreateStatusKamar)
	router.Get("/:id", GetStatusKamarById)
	router.Put("/:id", UpdateStatusKamar)
	router.Delete("/:id", DeleteStatusKamar)
}
