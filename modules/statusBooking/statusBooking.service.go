package statusbooking

import (
	"booking-hotel/databases"
	"booking-hotel/libs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateStatusBooking(c *fiber.Ctx) error {
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
	if err := databases.DB.Create(&statusBooking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Success Create Status Booking", 201)
}

func GetAllStatusBooking(c *fiber.Ctx) error {
	var statusBooking []StatusBooking
	if err := databases.DB.Find(&statusBooking).Error; err != nil {
		return libs.ResponseError(c, err, 500)
	}
	return libs.ResponseSuccess(c, statusBooking, 200)
}

func GetStatusBookingById(c *fiber.Ctx) error {
	statusBooking := StatusBooking{}
	id := c.Params("id")
	err := databases.DB.First(&statusBooking, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if err != nil {
		return libs.ResponseError(c, err, 500)
	}
	return libs.ResponseSuccess(c, statusBooking, 200)
}

func UpdateStatusBooking(c *fiber.Ctx) error {
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
	result := databases.DB.Model(&statusBooking).Where("status_booking_id = ?", id).Updates(&statusBooking)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Success Update Status Booking", 200)
}

func DeleteStatusBooking(c *fiber.Ctx) error {
	statusBooking := StatusBooking{}
	id := c.Params("id")
	result := databases.DB.Where("status_booking_id = ?", id).Delete(&statusBooking)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Success Delete Status Booking", 200)
}
