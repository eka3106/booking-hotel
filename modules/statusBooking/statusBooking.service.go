package statusbooking

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateStatusBooking(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	statusBooking := StatusBooking{}
	if err := c.BodyParser(&statusBooking); err != nil {
		return libs.ResponseError(c, err, 400)
	}
	if err := validate.Struct(statusBooking); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Table("status_booking").Create(&statusBooking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Success Create Status Booking", 201)
}

func GetAllStatusBooking(c *fiber.Ctx) error {
	var statusBooking []StatusBooking
	if err := databases.DB.Table("status_booking").Find(&statusBooking).Error; err != nil {
		return libs.ResponseError(c, err, 500)
	}
	return libs.ResponseSuccess(c, statusBooking, 200)
}

func GetStatusBookingById(c *fiber.Ctx) error {
	statusBooking := StatusBooking{}
	id := c.Params("id")
	err := databases.DB.Table("status_booking").First(&statusBooking, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, statusBooking, 200)
}

func UpdateStatusBooking(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	statusBooking := StatusBooking{}
	id := c.Params("id")
	if err := c.BodyParser(&statusBooking); err != nil {
		return libs.ResponseError(c, err, 400)
	}
	if err := validate.Struct(statusBooking); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	result := databases.DB.Table("status_booking").Model(&statusBooking).Where("status_booking_id = ?", id).Updates(&statusBooking)
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, "Success Update Status Booking", 200)
}

func DeleteStatusBooking(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	statusBooking := StatusBooking{}
	id := c.Params("id")
	result := databases.DB.Table("status_booking").Where("status_booking_id = ?", id).Delete(&statusBooking)
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, "Success Delete Status Booking", 200)
}
