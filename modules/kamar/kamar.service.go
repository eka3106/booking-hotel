package kamar

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	kamar := Kamar{}

	if err := c.BodyParser(&kamar); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(kamar); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Table("kamar").Create(&kamar).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create kamar", 201)
}

func GetAllKamar(c *fiber.Ctx) error {
	var kamar []Kamar
	if err := databases.DB.Table("kamar").Find(&kamar).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, kamar, 200)
}

func GetKamarById(c *fiber.Ctx) error {
	kamar := Kamar{}
	id := c.Params("id")
	err := databases.DB.Table("kamar").First(&kamar, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, kamar, 200)
}

func UpdateKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	kamar := Kamar{}
	id := c.Params("id")
	if err := databases.DB.Table("kamar").First(&kamar, id).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := c.BodyParser(&kamar); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(kamar); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	err := databases.DB.Table("kamar").Where("kamar_id = ?", id).Updates(&kamar)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success update kamar", 200)
}

func DeleteKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	kamar := Kamar{}
	id := c.Params("id")
	if err := databases.DB.Table("kamar").First(&kamar, id).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	err := databases.DB.Table("kamar").Delete(&kamar, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success delete kamar", 200)
}
