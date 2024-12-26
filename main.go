package main

import (
	_ "booking-hotel/config"
	"booking-hotel/databases"
	_ "booking-hotel/docs"
	"booking-hotel/middleware"
	"booking-hotel/modules/booking"
	"booking-hotel/modules/fasilitas"
	fasilitashotel "booking-hotel/modules/fasilitasHotel"
	hakakses "booking-hotel/modules/hakAkses"
	"booking-hotel/modules/hotel"
	"booking-hotel/modules/kamar"
	"booking-hotel/modules/pembayaran"
	statusbooking "booking-hotel/modules/statusBooking"
	statuskamar "booking-hotel/modules/statusKamar"
	tipekamar "booking-hotel/modules/tipeKamar"
	"booking-hotel/modules/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title Reservasi Hotel API
// @version 1.0
// @description API untuk reservasi hotel
// @description To access the API, you need to register and login first
// @BasePath /api/1.0
// @securityDefinitions.api_key Bearer:
// @in header
// @name Authorization
func main() {

	defer databases.SqlDb.Close()

	fiberApp := fiber.New()
	fiberApp.Use(cors.New())
	fiberApp.Get("/docs/*", swagger.HandlerDefault)
	fiberApp.Use(middleware.AuthJWT())
	booking.RouterBooking(fiberApp)
	fasilitas.RouterFasilitas(fiberApp)
	fasilitashotel.RouterFasilitasHotel(fiberApp)
	hakakses.RouterHakAkses(fiberApp)
	hotel.RouterHotel(fiberApp)
	kamar.RouterKamar(fiberApp)
	pembayaran.RouterPembayaran(fiberApp)
	statusbooking.RouterStatusBooking(fiberApp)
	statuskamar.RouterStatusKamar(fiberApp)
	tipekamar.RouterTipeKamar(fiberApp)
	user.RouterUser(fiberApp)

	fiberApp.Listen(":3000")
}
