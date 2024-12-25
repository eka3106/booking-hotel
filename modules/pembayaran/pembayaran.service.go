package pembayaran

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreatePembayaran(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	pembayaran := Pembayaran{}

	if err := c.BodyParser(&pembayaran); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}

	if err := validate.Struct(pembayaran); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}

	if err := databases.DB.Table("pembayaran").Create(&pembayaran).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create pembayaran", 201)
}

func GetAllPembayaran(c *fiber.Ctx) error {
	var pembayaran []Pembayaran
	if err := databases.DB.Table("pembayaran").Find(&pembayaran).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, pembayaran, 200)
}

func GetPembayaranById(c *fiber.Ctx) error {
	var pembayaran Pembayaran
	id := c.Params("id")
	err := databases.DB.Table("pembayaran").First(&pembayaran, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, pembayaran, 200)
}

func UpdatePembayaran(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	pembayaran := Pembayaran{}
	id := c.Params("id")
	if err := c.BodyParser(&pembayaran); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(pembayaran); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	err := databases.DB.Table("pembayaran").Where("pembayaran_id = ?", id).Updates(&pembayaran)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update pembayaran", 200)
}

func DeletePembayaran(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	id := c.Params("id")
	err := databases.DB.Table("pembayaran").Where("pembayaran_id = ?", id).Delete(Pembayaran{})
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete pembayaran", 200)
}
