package booking

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/kamar"
	"booking-hotel/modules/user"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

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
	if err := databases.DB.Table("kamar").Where("kamar_id = ?", booking.Kamar_id).First(&kamar); err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	} else if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Kamar not found", 404)
	}
	booking.User_id = claims.(*user.Claims).User_id
	booking.Hotel_id = kamar.Hotel_id
	booking.Total_biaya = kamar.Harga * booking.Jumlah_hari
	booking.Status_booking_id = 1
	var wg sync.WaitGroup
	updateKamar := make(chan bool)
	wg.Add(1)
	go func(ch chan bool) {
		if err := databases.DB.Table("kamar").Where("kamar_id = ?", booking.Kamar_id).Update("status_kamar_id", 3); err.Error != nil {
			ch <- false
		} else if err.RowsAffected == 0 {
			ch <- false
		} else {
			ch <- true
		}
	}(updateKamar)
	go func() {
		wg.Wait()
		close(updateKamar)
	}()
	isSuccess := <-updateKamar
	if !isSuccess {
		return libs.ResponseError(c, "Gagal booking", 400)
	}
	if err := databases.DB.Table("booking").Create(&booking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create booking", 201)
}

func GetAllBookingByUser(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	var booking []Booking
	if err := databases.DB.Preload("Status_booking").Table("booking").Where("user_id = ?", claims.(*user.Claims).User_id).Find(&booking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

func GetAllBookingByAdminHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	var booking []Booking
	if err := databases.DB.Preload("Status_booking").Table("booking").Where("hotel_id = ?", claims.(*user.Claims).Hotel_id).Find(&booking).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

func GetBookingUserById(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	booking := Booking{}
	id := c.Params("id")
	err := databases.DB.Preload("Status_booking").Table("booking").Where("user_id = ?", claims.(*user.Claims).User_id).First(&booking, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	if booking.User_id != c.Locals("user").(*user.Claims).User_id && c.Locals("user").(*user.Claims).Hak_akses_id != 1 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

func GetBookingHotelById(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	booking := Booking{}
	id := c.Params("id")
	err := databases.DB.Table("booking").First(&booking, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, booking, 200)
}

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
