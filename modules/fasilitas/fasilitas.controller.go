package fasilitas

import "github.com/gofiber/fiber/v2"

func RouterFasilitas(app *fiber.App) {
	router := app.Group("/api/1.0/fasilitas")
	router.Get("/", GetAllFasilitas)
	router.Post("/", CreateFasilitas)
	router.Get("/:id", GetFasilitasById)
	router.Put("/:id", UpdateFasilitas)
	router.Delete("/:id", DeleteFasilitas)
}
