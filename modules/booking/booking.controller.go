package booking

import "github.com/gofiber/fiber/v2"

func RouterBooking(app *fiber.App) {
	router := app.Group("/api/1.0")
	router.Post("/booking", CreateBooking)
	router.Get("/booking", GetAllBookingByUser)	
	router.Get("/booking/admin", GetAllBookingByAdminHotel)
	router.Get("/booking/detail/:id", GetBookingById)
	router.Patch("/booking/proses-reservasi", UpdateAdminBooking)

	// router.Delete("/booking/:id", CancelBooking)
}
