package pembayaran

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreatePembayaran(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	pembayaran := Pembayaran{}

	if err := c.BodyParser(&pembayaran); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}

	if err := validate.Struct(pembayaran); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}

	updateStatusBooking := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(ch chan bool) {
		defer wg.Done()
		if err := databases.DB.Table("booking").Where("booking_id = ?", pembayaran.Booking_id).Update("status_booking_id", 2); err.Error != nil {
			ch <- false
		} else if err.RowsAffected == 0 {
			ch <- false
		} else {
			ch <- true
		}
	}(updateStatusBooking)
	go func() {
		wg.Wait()
		close(updateStatusBooking)
	}()
	isSuccess := <-updateStatusBooking
	if !isSuccess {
		return libs.ResponseError(c, "Failed update status booking", 400)
	}
	if err := databases.DB.Table("pembayaran").Create(&pembayaran).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create pembayaran", 201)
}

func GetAllPembayaran(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	var pembayaran []ResponsePembayaran
	if err := databases.DB.Preload("Booking").Table("pembayaran").Joins("JOIN booking ON booking.booking_id = pembayaran.booking_id").Where("booking.user_id = ?", claims.(*user.Claims).User_id).Find(&pembayaran).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, pembayaran, 200)
}

func GetPembayaranById(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	var pembayaran ResponsePembayaran
	id := c.Params("id")
	err := databases.DB.Preload("Booking").Table("pembayaran").First(&pembayaran, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, pembayaran, 200)
}
