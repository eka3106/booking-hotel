package hakakses

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()


// CreateHakAkses godoc
// @Summary Create Hak Akses
// @Description Create Hak Akses
// @Tags Hak Akses
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreateHakAkses body RequestHakAkses true "Create Hak Akses"
// @Success 201 {string} string "Success Create Hak Akses"
// @Router /1.0/hak-akses [post]
func CreateHakAkses(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
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

// GetAllHakAkses godoc
// @Summary Get All Hak Akses
// @Description Get All Hak Akses
// @Tags Hak Akses
// @Accept json
// @Produce json
// @Success 200 {object} HakAkses
// @Router /1.0/hak-akses [get]
func GetAllHakAkses(c *fiber.Ctx) error {
	var hakAkses []HakAkses
	if err := databases.DB.Find(&hakAkses).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, hakAkses, 200)
}

// GetHakAksesById godoc
// @Summary Get Hak Akses By Id
// @Description Get Hak Akses By Id
// @Tags Hak Akses
// @Accept json
// @Produce json
// @Param id path int true "Hak Akses ID"
// @Success 200 {object} HakAkses
// @Router /1.0/hak-akses/{id} [get]
func GetHakAksesById(c *fiber.Ctx) error {
	hakAkses := HakAkses{}
	id := c.Params("id")
	err := databases.DB.First(&hakAkses, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, hakAkses, 200)
}

// UpdateHakAkses godoc
// @Summary Update Hak Akses
// @Description Update Hak Akses
// @Tags Hak Akses
// @Accept json
// @Produce json
// @Param id path int true "Hak Akses ID"
// @Param Authorization header string true "Bearer"
// @Param UpdateHakAkses body RequestHakAkses true "Update Hak Akses"	
// @Success 200 {string} string "Success Update Hak Akses"
// @Router /1.0/hak-akses/{id} [put]
func UpdateHakAkses(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
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
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 400)
	}
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update hak akses", 200)
}

// DeleteHakAkses godoc
// @Summary Delete Hak Akses
// @Description Delete Hak Akses
// @Tags Hak Akses
// @Accept json
// @Produce json
// @Param id path int true "Hak Akses ID"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} string "Success Delete Hak Akses"
// @Router /1.0/hak-akses/{id} [delete]
func DeleteHakAkses(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	hakAkses := HakAkses{}
	id := c.Params("id")
	result := databases.DB.Where("hak_akses_id = ?", id).Delete(&hakAkses)
	if result.Error != nil {
		return libs.ResponseError(c, result.Error.Error(), 400)
	}
	if result.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete hak akses", 200)
}
