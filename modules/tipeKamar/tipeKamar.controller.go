package tipekamar

import "github.com/gofiber/fiber/v2"

func RouterTipeKamar(app *fiber.App) {
	router := app.Group("/api/1.0/tipe-kamar")
	router.Get("/", GetAllTypeKamar)
	router.Post("/", CreateTypeKamar)
	router.Get("/:id", GetTypeKamarById)
	router.Put("/:id", UpdateTypeKamar)
	router.Delete("/:id", DeleteTypeKamar)
}
