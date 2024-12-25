package main

import (
	_ "booking-hotel/config"
	"booking-hotel/databases"
	"booking-hotel/middleware"
	"booking-hotel/modules/booking"
	"booking-hotel/modules/fasilitas"
	fasilitashotel "booking-hotel/modules/fasilitasHotel"
	hakakses "booking-hotel/modules/hakAkses"
	"booking-hotel/modules/hotel"
	"booking-hotel/modules/kamar"
	"booking-hotel/modules/pembayaran"
	"booking-hotel/modules/rating"
	statusbooking "booking-hotel/modules/statusBooking"
	statuskamar "booking-hotel/modules/statusKamar"
	tipekamar "booking-hotel/modules/tipeKamar"
	"booking-hotel/modules/user"

	"github.com/gofiber/fiber/v2"
)

func main() {

	defer databases.SqlDb.Close()

	fiberApp := fiber.New()

	fiberApp.Use(middleware.AuthJWT())

	booking.RouterBooking(fiberApp)
	fasilitas.RouterFasilitas(fiberApp)
	fasilitashotel.RouterFasilitasHotel(fiberApp)
	hakakses.RouterHakAkses(fiberApp)
	hotel.RouterHotel(fiberApp)
	kamar.RouterKamar(fiberApp)
	pembayaran.RouterPembayaran(fiberApp)
	rating.RouterRating(fiberApp)
	statusbooking.RouterStatusBooking(fiberApp)
	statuskamar.RouterStatusKamar(fiberApp)
	tipekamar.RouterTipeKamar(fiberApp)
	user.RouterUser(fiberApp)

	fiberApp.Listen(":3000")
}
