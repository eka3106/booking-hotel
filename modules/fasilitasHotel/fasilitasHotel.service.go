package fasilitashotel

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitasHotel := FasilitasHotel{}
	if err := c.BodyParser(&fasilitasHotel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err := validate.Struct(fasilitasHotel)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Table("fasilitas_hotel").Create(&fasilitasHotel); err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create fasilitas hotel", 201)
}

func GetAllFasilitasHotel(c *fiber.Ctx) error {
	var fasilitasHotel []FasilitasHotel
	if err := databases.DB.Table("fasilitas_hotel").Find(&fasilitasHotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}

func GetFasilitasHotelById(c *fiber.Ctx) error {
	fasilitasHotel := FasilitasHotel{}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas_hotel").First(&fasilitasHotel, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}

func UpdateFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitasHotel := FasilitasHotel{}
	if err := c.BodyParser(&fasilitasHotel); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas_hotel").Where("fasilitas_hotel_id = ?", id).Updates(&fasilitasHotel)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update fasilitas hotel", 200)
}

func DeleteFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas_hotel").Where("fasilitas_hotel_id = ?", id).Delete(&FasilitasHotel{})
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete fasilitas hotel", 200)
}
