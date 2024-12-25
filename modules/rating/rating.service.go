package rating

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateRating(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	rating := Rating{}
	if err := c.BodyParser(&rating); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	rating.User_id = claims.(*user.Claims).IdUser
	if err := validate.Struct(rating); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if err := databases.DB.Table("rating").Create(&rating).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create rating", 201)
}

func GetAllRating(c *fiber.Ctx) error {
	var rating []Rating
	if err := databases.DB.Table("rating").Find(&rating).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, rating, 200)
}

func GetRatingById(c *fiber.Ctx) error {
	rating := Rating{}
	id := c.Params("id")
	err := databases.DB.Table("rating").First(&rating, id)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, rating, 200)
}

func UpdateRating(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	rating := RequestRating{}
	id := c.Params("id")
	if err := c.BodyParser(&rating); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(rating); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	result := databases.DB.Table("rating").Where("rating_id = ? AND user_id = ?", id, claims.(*user.Claims).IdUser).Updates(&rating)
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success update rating", 200)
}

func DeleteRating(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	id := c.Params("id")
	err := databases.DB.Table("rating").Where("rating_id = ? AND user_id = ?", id, claims.(*user.Claims).IdUser).Delete(&Rating{})
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success delete rating", 200)
}

func GetRatingHotelId(c *fiber.Ctx) error {
	var rating ResponseRating
	var ratings []Rating
	id := c.Params("id")
	err := databases.DB.Table("rating").Where("hotel_id = ?", id).Find(&ratings)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	var totalRating int
	for _, value := range ratings {
		totalRating += value.Rating
	}
	rataRata := totalRating / len(ratings)
	rating.Rating = rataRata
	rating.Hotel_id, _ = strconv.Atoi(id)
	return libs.ResponseSuccess(c, rating, 200)
}

func GetRatingUserId(c *fiber.Ctx) error {
	var rating []Rating
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	id := claims.(*user.Claims).IdUser
	err := databases.DB.Table("rating").Where("user_id = ?", id).Find(&rating)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, rating, 200)
}

func GetRatingByHotelIdAndUserId(c *fiber.Ctx) error {
	var rating Rating
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	id := c.Params("id")
	hotelId, _ := strconv.Atoi(id)
	userId := claims.(*user.Claims).IdUser
	err := databases.DB.Table("rating").Where("hotel_id = ? AND user_id = ?", hotelId, userId).First(&rating)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, rating, 200)
}
