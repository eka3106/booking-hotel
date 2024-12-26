package pembayaran

import "github.com/gofiber/fiber/v2"

func RouterPembayaran(app *fiber.App) {
	router := app.Group("/api/1.0/pembayaran")
	router.Post("/", CreatePembayaran)
	router.Get("/", GetAllPembayaran)
	router.Get("/:id", GetPembayaranById)
}
