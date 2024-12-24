package hakakses

import (
	"booking-hotel/databases"
	"booking-hotel/libs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateHakAkses(c *fiber.Ctx) error {
	hakAkses := HakAkses{}
	if err := c.BodyParser(&hakAkses); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	err := validate.Struct(hakAkses)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Create(&hakAkses).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create hak akses", 201)
}

func GetAllHakAkses(c *fiber.Ctx) error {
	var hakAkses []HakAkses
	if err := databases.DB.Find(&hakAkses).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, hakAkses, 200)
}

func GetHakAksesById(c *fiber.Ctx) error {
	hakAkses := HakAkses{}
	id := c.Params("id")
	err := databases.DB.First(&hakAkses, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, hakAkses, 200)
}

func UpdateHakAkses(c *fiber.Ctx) error {
	hakAkses := HakAkses{}
	id := c.Params("id")
	if err := c.BodyParser(&hakAkses); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(hakAkses); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	result := databases.DB.Model(&hakAkses).Where("hak_akses_id = ?", id).Updates(&hakAkses)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success update hak akses", 200)
}

func DeleteHakAkses(c *fiber.Ctx) error {
	hakAkses := HakAkses{}
	id := c.Params("id")
	result := databases.DB.Where("hak_akses_id = ?", id).Delete(&hakAkses)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success delete hak akses", 200)
}
