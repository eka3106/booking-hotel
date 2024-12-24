package hakakses

import "github.com/gofiber/fiber/v2"

func RouterHakAkses(app *fiber.App) {
	router := app.Group("/api/1.0/hak-akses")
	router.Get("/", GetAllHakAkses)
	router.Post("/", CreateHakAkses)
	router.Get("/:id", GetHakAksesById)
	router.Put("/:id", UpdateHakAkses)
	router.Delete("/:id", DeleteHakAkses)
}
