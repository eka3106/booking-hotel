package main

import (
	_ "booking-hotel/config"
	"booking-hotel/databases"
	tipekamar "booking-hotel/modules/tipeKamar"

	"github.com/gofiber/fiber/v2"
)

func main() {

	defer databases.SqlDb.Close()

	fiberApp := fiber.New()

	tipekamar.RouterTipeKamar(fiberApp)

	fiberApp.Listen(":3000")
}
