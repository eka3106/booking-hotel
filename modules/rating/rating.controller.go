package rating

import "github.com/gofiber/fiber/v2"

func RouterRating(app *fiber.App) {
	router := app.Group("/api/v1")
	router.Post("/rating", CreateRating)
	router.Get("/rating", GetAllRating)
	router.Get("/rating/:id", GetRatingById)
	router.Put("/rating/:id", UpdateRating)
	router.Get("/rating/by-user", GetRatingUserId)
	router.Get("/rating/one-hotel/:id", GetRatingByHotelIdAndUserId)
	router.Get("/rating/by-hotel/:id", GetRatingHotelId)
	router.Delete("/rating/:id", DeleteRating)
}
