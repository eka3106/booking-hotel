package booking

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/kamar"
	"booking-hotel/modules/user"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var validate = validator.New()

// CreateBooking godoc
// @Summary Create Booking
// @Description Create Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreateBooking body RequestCreateBookingSwagger true "Create Booking"
// @Success 201 {string} string "Success Create Booking"
// @Router /1.0/booking [post]
func CreateBooking(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	booking := RequestCreateBooking{}
	if err := c.BodyParser(&booking); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(booking); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	kamar := kamar.Kamar{}
	if err := databases.DB.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Table("kamar").Where("kamar_id = ? AND status_kamar_id = 1", booking.Kamar_id).Clauses(clause.Locking{Strength:"UPDATE"}).First(&kamar).Error; err != nil {
				return err
			} 
			booking.User_id = claims.(*user.Claims).User_id
			booking.Hotel_id = kamar.Hotel_id
			booking.Total_biaya = kamar.Harga * booking.Jumlah_hari
			booking.Status_booking_id = 1
			if err := tx.Table("kamar").Where("kamar_id = ? AND status_kamar_id = 1", booking.Kamar_id).Update("status_kamar_id", 3).Error; err != nil {
				return err
			} 
			if err := tx.Table("booking").Create(&booking).Error; err != nil {
				return err
			}
			return nil
		},
	); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	} 
	return libs.ResponseSuccess(c, "Success create booking", 201)
}

// GetAllBookingByUser godoc
// @Summary Get All Booking
// @Description Get All Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} Booking
// @Router /1.0/booking [get]
func GetAllBookingByUser(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	var booking []ResponseBookingUser
	if err := databases.DB.Preload("Kamar").Preload("Hotel").Preload("Status_booking").Joins("JOIN kamar ON kamar.kamar_id = booking.kamar_id").Joins("JOIN hotel ON hotel.hotel_id = booking.hotel_id").Joins("JOIN status_booking ON status_booking.status_booking_id = booking.status_booking_id").Table("booking").Where("user_id = ?", claims.(*user.Claims).User_id).Find(&booking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}	
	return libs.ResponseSuccess(c, booking, 200)
}

// GetAllBookingByAdminHotel godoc
// @Summary Get All Booking on Hotel
// @Description Get All Booking on Hotel harus admin hotel
// @Tags Booking
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} Booking
// @Router /1.0/booking/admin [get]
func GetAllBookingByAdminHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	var booking []ResponseBookingAdmin
	if err := databases.DB.Preload("User").Preload("Kamar").Preload("Hotel").Preload("Status_booking").
	Joins("JOIN users ON users.user_id = booking.user_id").
	Joins("JOIN kamar ON kamar.kamar_id = booking.kamar_id").
	Joins("JOIN hotel ON hotel.hotel_id = booking.hotel_id").
	Joins("JOIN status_booking ON status_booking.status_booking_id = booking.status_booking_id").
	Table("booking").Where("booking.hotel_id = ?", claims.(*user.Claims).Hotel_id).Find(&booking).Error; err != nil {
	return libs.ResponseError(c, err.Error(), 400)
}
	return libs.ResponseSuccess(c, booking, 200)
}

// GetBookingById godoc
// @Summary Get Booking By Id
// @Description Get Booking By Id
// @Tags Booking
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param Authorization header string true "Bearer"
// @Success 200 {object} Booking
// @Router /1.0/booking/{id} [get]
func GetBookingById(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	booking := ResponseBookingAdmin{}
	id := c.Params("id")
	err := databases.DB.Preload("User").Preload("Kamar").Preload("Hotel").Preload("Status_booking").
	Joins("JOIN users ON users.user_id = booking.user_id").
	Joins("JOIN kamar ON kamar.kamar_id = booking.kamar_id").
	Joins("JOIN hotel ON hotel.hotel_id = booking.hotel_id").
	Joins("JOIN status_booking ON status_booking.status_booking_id = booking.status_booking_id").
	Table("booking").First(&booking, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if booking.User_id != c.Locals("user").(*user.Claims).User_id && c.Locals("user").(*user.Claims).Hotel_id != booking.Hotel_id {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

// UpdateAdminBooking godoc
// @Summary Update Booking
// @Description Update Booking
// @Tags Booking
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param UpdateBooking body RequestUpdateBooking true "Update Booking"
// @Success 200 {string} string "Success Update Booking"
// @Router /1.0/booking/proses-reservasi [patch]
func UpdateAdminBooking(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	booking := RequestAdmin{}
	if err := c.BodyParser(&booking); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	booking.Tanggal_check_out = libs.GetTimeNow()
	if err := validate.Struct(booking); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	if booking.Status_booking_id == 3 {
		booking.Tanggal_check_in = libs.GetTimeNow()
		booking.Tanggal_check_out = ""
		booking.Status_kamar_id = 2
	}
	if booking.Status_booking_id == 4 {
		booking.Tanggal_check_out = libs.GetTimeNow()
		booking.Status_kamar_id = 5
	}
	if booking.Status_booking_id == 5 {
		booking.Status_kamar_id = 1
	}
	if booking.Status_booking_id == 6 {
		booking.Status_kamar_id = 1
	}
	var wg sync.WaitGroup
	updateKamar := make(chan bool)
	wg.Add(1)
	go func(ch chan bool, statusKamar int) {
		if err := databases.DB.Table("kamar").Where("kamar_id = ?", booking.Kamar_id).Update("status_kamar_id", statusKamar); err.Error != nil {
			ch <- false
		} else if err.RowsAffected == 0 {
			ch <- false
		} else {
			ch <- true
		}
		defer wg.Done()
	}(updateKamar, booking.Status_kamar_id)
	go func() {
		wg.Wait()
		close(updateKamar)
	}()
	isSuccess := <-updateKamar
	if !isSuccess {
		return libs.ResponseError(c, "Gagal update, coba lagi", 400)
	}
	insertBooking := InsertBooking{
		Booking_id:        booking.Booking_id,
		Kamar_id:          booking.Kamar_id,
		Tanggal_check_in:  booking.Tanggal_check_in,
		Tanggal_check_out: booking.Tanggal_check_out,
		Status_booking_id: booking.Status_booking_id,
	}
	err := databases.DB.Table("booking").Where("booking_id = ?", booking.Booking_id).Updates(insertBooking)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update booking", 200)
}

// func DeleteBooking(c *fiber.Ctx) error {
// 	claims := c.Locals("user")
// 	if claims == nil {
// 		return libs.ResponseError(c, "Unauthorized", 401)
// 	}
// 	id := c.Params("id")
// 	booking := Booking{}
// 	err := databases.DB.Table("booking").Where("booking_id = ?", id).Delete(&booking)
// 	if err.Error != nil {
// 		return libs.ResponseError(c, err.Error.Error(), 400)
// 	}
// 	if err.RowsAffected == 0 {
// 		return libs.ResponseError(c, "Data not found", 404)
// 	}
// 	return libs.ResponseSuccess(c, "Success delete booking", 200)
// }
