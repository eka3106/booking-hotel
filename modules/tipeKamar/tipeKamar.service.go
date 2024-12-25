package tipekamar

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateTipeKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	TipeKamar := TipeKamar{}
	if err := c.BodyParser(&TipeKamar); err != nil {
		return c.Status(400).JSON(err)
	}
	err := validate.Struct(TipeKamar)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}

	result := databases.DB.Table("tipe_kamar").Create(&TipeKamar)
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Succes Create Tipe Kamar", 201)
}

func GetAllTipeKamar(c *fiber.Ctx) error {
	var TipeKamar []TipeKamar
	result := databases.DB.Table("tipe_kamar").Find(&TipeKamar)
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, TipeKamar, 200)
}

func GetTipeKamarById(c *fiber.Ctx) error {
	TipeKamar := TipeKamar{}
	id := c.Params("id")
	result := databases.DB.Table("tipe_kamar").First(&TipeKamar, id)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, TipeKamar, 200)
}

func UpdateTipeKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	TipeKamar := TipeKamar{}
	id := c.Params("id")
	if err := c.BodyParser(&TipeKamar); err != nil {
		return libs.ResponseError(c, err, 400)
	}
	err := validate.Struct(TipeKamar)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	result := databases.DB.Table("tipe_kamar").Where("tipe_kamar_id = ?", id).Updates(&TipeKamar)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Succes Update Tipe Kamar", 200)
}

func DeleteTipeKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	id := c.Params("id")
	result := databases.DB.Table("tipe_kamar").Delete(&TipeKamar{}, id)
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 500)
	}
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data Not Found", 404)
	}
	return libs.ResponseSuccess(c, "Succes Delete Tipe Kamar", 200)
}
