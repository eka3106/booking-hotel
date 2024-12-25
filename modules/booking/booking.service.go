package booking

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateBooking(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	booking := Booking{}
	if err := c.BodyParser(&booking); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	booking.User_id = claims.(*user.Claims).User_id
	if err := validate.Struct(booking); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Table("booking").Create(&booking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create booking", 201)
}

func GetAllBooking(c *fiber.Ctx) error {
	var booking []Booking
	if err := databases.DB.Table("booking").Find(&booking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

func GetBookingById(c *fiber.Ctx) error {
	booking := Booking{}
	id := c.Params("id")
	err := databases.DB.Table("booking").First(&booking, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

func GetBookingByUserId(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	booking := []Booking{}
	err := databases.DB.Table("booking").Where("user_id = ?", claims.(*user.Claims).User_id).Find(&booking)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

func GetBookingByHotelId(c *fiber.Ctx) error {
	booking := []Booking{}
	id := c.Params("id")
	err := databases.DB.Table("booking").Where("hotel_id = ?", id).Find(&booking)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

func UpdateBookingCheckIn(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	booking := RequestCheckIn{}
	if err := c.BodyParser(&booking); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	booking.Tanggal_check_in = libs.GetTimeNow()
	if err := validate.Struct(booking); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	err := databases.DB.Table("booking").Where("booking_id = ?", booking.Booking_id).Updates(&booking)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update booking", 200)
}

func UpdateBookingCheckOut(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	booking := RequestCheckOut{}
	if err := c.BodyParser(&booking); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	booking.Tanggal_check_out = libs.GetTimeNow()
	if err := validate.Struct(booking); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	err := databases.DB.Table("booking").Where("booking_id = ?", booking.Booking_id).Updates(&booking)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update booking", 200)
}

func DeleteBooking(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	booking := Booking{}
	id := c.Params("id")
	err := databases.DB.Table("booking").Delete(&booking, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete booking", 200)
}
