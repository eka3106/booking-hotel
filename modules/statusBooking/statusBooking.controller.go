package statusbooking

import "github.com/gofiber/fiber/v2"

func RouterStatusBooking(app *fiber.App) {
	router := app.Group("/api/1.0/status-booking")
	router.Get("/", GetAllStatusBooking)
	router.Post("/", CreateStatusBooking)
	router.Get("/:id", GetStatusBookingById)
	router.Put("/:id", UpdateStatusBooking)
	router.Delete("/:id", DeleteStatusBooking)
}
