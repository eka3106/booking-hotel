package statusbooking

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CreateStatusBooking godoc
// @Summary Create Status Booking
// @Description Create Status Booking
// @Tags Status Booking
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreateStatusBooking body RequestStatusBooking true "Create Status Booking"
// @Success 201 {string} string "Success Create Status Booking"
// @Router /1.0/status-booking [post]
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

// GetAllStatusBooking godoc
// @Summary Get All Status Booking
// @Description Get All Status Booking
// @Tags Status Booking
// @Accept json
// @Produce json
// @Success 200 {object} StatusBooking
// @Router /1.0/status-booking [get]
func GetAllStatusBooking(c *fiber.Ctx) error {
	var statusBooking []StatusBooking
	if err := databases.DB.Table("status_booking").Find(&statusBooking).Error; err != nil {
		return libs.ResponseError(c, err, 500)
	}
	return libs.ResponseSuccess(c, statusBooking, 200)
}

// GetStatusBookingById godoc
// @Summary Get Status Booking By ID
// @Description Get Status Booking By ID
// @Tags Status Booking
// @Accept json
// @Produce json
// @Param id path int true "ID Status Booking"
// @Success 200 {object} StatusBooking
// @Router /1.0/status-booking/{id} [get]
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

// UpdateStatusBooking godoc
// @Summary Update Status Booking
// @Description Update Status Booking
// @Tags Status Booking
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID Status Booking"
// @Param UpdateStatusBooking body RequestStatusBooking true "Update Status Booking"
// @Success 200 {string} string "Success Update Status Booking"
// @Router /1.0/status-booking/{id} [put]
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

// DeleteStatusBooking godoc
// @Summary Delete Status Booking
// @Description Delete Status Booking
// @Tags Status Booking
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID Status Booking"
// @Success 200 {string} string "Success Delete Status Booking"
// @Router /1.0/status-booking/{id} [delete]
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
