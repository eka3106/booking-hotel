package hotel

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	hotel := Hotel{}

	if err := c.BodyParser(&hotel); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}

	if err := validate.Struct(hotel); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}

	if err := databases.DB.Table("hotel").Create(&hotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}

	return libs.ResponseSuccess(c, "Success create hotel", 201)
}

func GetAllHotel(c *fiber.Ctx) error {
	var hotel []Hotel
	if err := databases.DB.Table("hotel").Find(&hotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, hotel, 200)
}

func GetHotelById(c *fiber.Ctx) error {
	hotel := Hotel{}
	id := c.Params("id")
	err := databases.DB.Table("hotel").First(&hotel, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, hotel, 200)
}

func UpdateHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	hotel := Hotel{}
	id := c.Params("id")
	if err := c.BodyParser(&hotel); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(hotel); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	err := databases.DB.Table("hotel").Where("hotel_id = ?", id).Updates(&hotel)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success update hotel", 200)
}

func DeleteHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	hotel := Hotel{}
	id := c.Params("id")
	err := databases.DB.Table("hotel").Delete(&hotel, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success delete hotel", 200)
}
