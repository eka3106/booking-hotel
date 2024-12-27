package hotel

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CreateHotel godoc
// @Summary Create Hotel
// @Description Create Hotel
// @Tags Hotel
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreateHotel body RequestHotel true "Create Hotel"
// @Success 201 {string} string "Success Create Hotel"
// @Router /1.0/hotel [post]
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

// GetAllHotel godoc
// @Summary Get All Hotel
// @Description Get All Hotel
// @Tags Hotel
// @Accept json
// @Produce json
// @Success 200 {object} Hotel
// @Router /1.0/hotel [get]
func GetAllHotel(c *fiber.Ctx) error {
	var hotel []Hotel
	if err := databases.DB.Table("hotel").Find(&hotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, hotel, 200)
}

// GetHotelById godoc
// @Summary Get Hotel By ID
// @Description Get Hotel By ID
// @Tags Hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {object} Hotel
// @Router /1.0/hotel/{id} [get]
func GetHotelById(c *fiber.Ctx) error {
	hotel := Hotel{}
	id := c.Params("id")
	err := databases.DB.Table("hotel").First(&hotel, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, hotel, 200)
}

// UpdateHotel godoc
// @Summary Update Hotel
// @Description Update Hotel
// @Tags Hotel
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "Hotel ID"
// @Param UpdateHotel body RequestHotel true "Update Hotel"
// @Success 200 {string} string "Success Update Hotel"
// @Router /1.0/hotel/{id} [put]
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
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update hotel", 200)
}

// DeleteHotel godoc
// @Summary Delete Hotel
// @Description Delete Hotel
// @Tags Hotel
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "Hotel ID"
// @Success 200 {string} string "Success Delete Hotel"
// @Router /1.0/hotel/{id} [delete]
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
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete hotel", 200)
}
