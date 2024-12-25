package fasilitas

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateFasilitas(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitas := Fasilitas{}
	if err := c.BodyParser(&fasilitas); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	err := validate.Struct(fasilitas)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Table("fasilitas").Create(&fasilitas).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create fasilitas", 201)
}

func GetAllFasilitas(c *fiber.Ctx) error {
	var fasilitas []Fasilitas
	if err := databases.DB.Table("fasilitas").Find(&fasilitas).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, fasilitas, 200)
}

func GetFasilitasById(c *fiber.Ctx) error {
	fasilitas := Fasilitas{}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas").First(&fasilitas, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, fasilitas, 200)
}

func UpdateFasilitas(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitas := Fasilitas{}
	id := c.Params("id")
	if err := c.BodyParser(&fasilitas); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(fasilitas); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	err := databases.DB.Table("fasilitas").Where("fasilitas_id = ?", id).Updates(&fasilitas)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update fasilitas", 200)
}

func DeleteFasilitas(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitas := Fasilitas{}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas").Delete(&fasilitas, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete fasilitas", 200)
}
