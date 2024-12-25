package pembayaran

import "github.com/gofiber/fiber/v2"

func RouterPembayaran(app *fiber.App) {
	router := app.Group("/api/v1/pembayaran")
	router.Post("/", CreatePembayaran)
	router.Get("/", GetAllPembayaran)
	router.Get("/:id", GetPembayaranById)
	router.Put("/:id", UpdatePembayaran)
	router.Delete("/:id", DeletePembayaran)
}
