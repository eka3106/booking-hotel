package pembayaran

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CreatePembayaran godoc
// @Summary Create Pembayaran
// @Description Create Pembayaran
// @Tags Pembayaran
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreatePembayaran body RequestPembayaran true "Create Pembayaran"
// @Success 201 {string} string "Success Create Pembayaran"
// @Router /1.0/pembayaran [post]
func CreatePembayaran(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
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

	updateStatusBooking := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(ch chan bool) {
		defer wg.Done()
		if err := databases.DB.Table("booking").Where("booking_id = ?", pembayaran.Booking_id).Update("status_booking_id", 2); err.Error != nil {
			ch <- false
		} else if err.RowsAffected == 0 {
			ch <- false
		} else {
			ch <- true
		}
	}(updateStatusBooking)
	go func() {
		wg.Wait()
		close(updateStatusBooking)
	}()
	isSuccess := <-updateStatusBooking
	if !isSuccess {
		return libs.ResponseError(c, "Failed update status booking", 400)
	}
	if err := databases.DB.Table("pembayaran").Create(&pembayaran).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create pembayaran", 201)
}

// GetAllPembayaran godoc
// @Summary Get All Pembayaran
// @Description Get All Pembayaran
// @Tags Pembayaran
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} ResponsePembayaran
// @Router /1.0/pembayaran [get]
func GetAllPembayaran(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	var pembayaran []ResponsePembayaran
	if err := databases.DB.Preload("Booking").Preload("Booking.Status_booking").
        Joins("JOIN booking ON booking.booking_id = pembayaran.booking_id").
        Joins("JOIN status_booking ON booking.status_booking_id = status_booking.status_booking_id").Table("pembayaran").
        Where("booking.user_id = ?", claims.(*user.Claims).User_id).
        Find(&pembayaran).Error; err != nil {
        return libs.ResponseError(c, err.Error(), 400)
    }
	
	return libs.ResponseSuccess(c, pembayaran, 200)
}

// GetPembayaranById godoc
// @Summary Get Pembayaran By ID
// @Description Get Pembayaran By ID
// @Tags Pembayaran
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param id path int true "ID Pembayaran"
// @Success 200 {object} ResponsePembayaran
// @Router /1.0/pembayaran/{id} [get]
func GetPembayaranById(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	var pembayaran ResponsePembayaran
	id := c.Params("id")

	err := databases.DB.Preload("Booking").Preload("Booking.Status_booking").
        Joins("JOIN booking ON booking.booking_id = pembayaran.booking_id").
        Joins("JOIN status_booking ON booking.status_booking_id = status_booking.status_booking_id").Table("pembayaran").
        Where("pembayaran.pembayaran_id = ?", id).
        Find(&pembayaran)    
    
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, pembayaran, 200)
}
