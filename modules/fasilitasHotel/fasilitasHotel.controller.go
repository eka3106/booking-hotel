package fasilitashotel

import "github.com/gofiber/fiber/v2"

func RouterFasilitasHotel(app *fiber.App) {
	router := app.Group("/api/1.0/fasilitas-hotel")
	router.Post("/", CreateFasilitasHotel)
	router.Get("/", GetAllFasilitasHotel)
	router.Get("/:id", GetFasilitasHotelById)
	router.Put("/:id", UpdateFasilitasHotel)
	router.Get("/by-hotel/:id", GetFasilitasByHotelId)
	router.Get("/by-fasilitas/:id", GetFasilitasByFasilitasId)
	router.Delete("/:id", DeleteFasilitasHotel)
}
