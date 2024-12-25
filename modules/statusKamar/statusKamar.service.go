package statuskamar

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

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
	if err := databases.DB.Create(&statusKamar).Error; err != nil {
		return libs.ResponseError(c, err, 500)
	}
	return libs.ResponseSuccess(c, "Success Create Status Kamar", 201)
}

func GetAllStatusKamar(c *fiber.Ctx) error {
	var statusKamar []StatusKamar
	if err := databases.DB.Find(&statusKamar).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 500)
	}
	if len(statusKamar) == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, statusKamar, 200)
}

func GetStatusKamarById(c *fiber.Ctx) error {
	statusKamar := StatusKamar{}
	id := c.Params("id")
	err := databases.DB.First(&statusKamar, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, statusKamar, 200)
}

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
	result := databases.DB.Model(&statusKamar).Where("status_kamar_id = ?", id).Updates(&statusKamar)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Success Update Status Kamar", 200)
}

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
	err := databases.DB.Delete(&statusKamar, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error, 500)
	}
	return libs.ResponseSuccess(c, "Success Delete Status Kamar", 200)
}
