package tipekamar

import "github.com/gofiber/fiber/v2"

func RouterTipeKamar(app *fiber.App) {
	router := app.Group("/api/1.0/tipe-kamar")
	router.Get("/", GetAllTipeKamar)
	router.Post("/", CreateTipeKamar)
	router.Get("/:id", GetTipeKamarById)
	router.Put("/:id", UpdateTipeKamar)
	router.Delete("/:id", DeleteTipeKamar)
}
