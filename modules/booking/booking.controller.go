package booking

import "github.com/gofiber/fiber/v2"

func RouterBooking(app *fiber.App) {
	router := app.Group("/api/v1")
	router.Post("/booking", CreateBooking)
	router.Get("/booking", GetAllBooking)
	router.Get("/booking/:id", GetBookingById)
	router.Get("/booking/by-user", GetBookingByUserId)
	router.Get("/booking/by-hotel/:id", GetBookingByHotelId)
	router.Patch("/booking/check-in", UpdateBookingCheckIn)
	router.Patch("booking/check-out", UpdateBookingCheckOut)
	router.Delete("/:id", DeleteBooking)
}
