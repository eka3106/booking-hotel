package statuskamar

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CreateStatusKamar godoc
// @Summary Create Status Kamar
// @Description Create Status Kamar
// @Tags Status Kamar
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreateStatusKamar body RequestStatusKamar true "Create Status Kamar"
// @Success 201 {string} string "Success Create Status Kamar"
// @Router /1.0/status-kamar [post]
func CreateStatusKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	statusKamar := StatusKamar{}
	if err := c.BodyParser(&statusKamar); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(statusKamar); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Table("status_kamar").Create(&statusKamar).Error; err != nil {
		return libs.ResponseError(c, err, 500)
	}
	return libs.ResponseSuccess(c, "Success Create Status Kamar", 201)
}

// GetAllStatusKamar godoc
// @Summary Get All Status Kamar
// @Description Get All Status Kamar
// @Tags Status Kamar
// @Accept json
// @Produce json
// @Success 200 {object} StatusKamar
// @Router /1.0/status-kamar [get]
func GetAllStatusKamar(c *fiber.Ctx) error {
	var statusKamar []StatusKamar
	if err := databases.DB.Table("status_kamar").Find(&statusKamar).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 500)
	}
	if len(statusKamar) == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, statusKamar, 200)
}

// GetStatusKamarById godoc
// @Summary Get Status Kamar By ID
// @Description Get Status Kamar By ID
// @Tags Status Kamar
// @Accept json
// @Produce json
// @Param id path int true "Status Kamar ID"
// @Success 200 {object} StatusKamar
// @Router /1.0/status-kamar/{id} [get]
func GetStatusKamarById(c *fiber.Ctx) error {
	statusKamar := StatusKamar{}
	id := c.Params("id")
	err := databases.DB.Table("status_kamar").First(&statusKamar, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, statusKamar, 200)
}

// UpdateStatusKamar godoc
// @Summary Update Status Kamar
// @Description Update Status Kamar
// @Tags Status Kamar
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "Status Kamar ID"
// @Param UpdateStatusKamar body RequestStatusKamar true "Update Status Kamar"
// @Success 200 {string} string "Success Update Status Kamar"
// @Router /1.0/status-kamar/{id} [put]
func UpdateStatusKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	statusKamar := StatusKamar{}
	id := c.Params("id")
	if err := c.BodyParser(&statusKamar); err != nil {
		return libs.ResponseError(c, err, 400)
	}
	err := validate.Struct(statusKamar)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	result := databases.DB.Table("status_kamar").Model(&statusKamar).Where("status_kamar_id = ?", id).Updates(&statusKamar)
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, "Success Update Status Kamar", 200)
}

// DeleteStatusKamar godoc
// @Summary Delete Status Kamar
// @Description Delete Status Kamar
// @Tags Status Kamar
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "Status Kamar ID"
// @Success 200 {string} string "Success Delete Status Kamar"
// @Router /1.0/status-kamar/{id} [delete]
func DeleteStatusKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	statusKamar := StatusKamar{}
	id := c.Params("id")
	err := databases.DB.Table("status_kamar").Delete(&statusKamar, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error, 500)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, "Success Delete Status Kamar", 200)
}
