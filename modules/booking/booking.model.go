package booking

import (
	"booking-hotel/modules/hotel"
	"booking-hotel/modules/kamar"
	statusbooking "booking-hotel/modules/statusBooking"
	"booking-hotel/modules/user"
	"time"
)

type Booking struct {
	Booking_id        int                         `json:"booking_id" gorm:"primaryKey"`
	Kamar_id          int                         `json:"kamar_id" gorm:"foreignKey:kamar_id" validate:"required"`
	Hotel_id          int                         `json:"hotel_id" gorm:"foreignKey:hotel_id" validate:"required"`
	User_id           int                         `json:"user_id" gorm:"foreignKey:user_id" validate:"required"`
	Tanggal_check_in  string                      `json:"tanggal_check_in" gorm:"type:date" `
	Jumlah_hari       int                         `json:"jumlah_hari" gorm:"type:int"`
	Tanggal_check_out string                      `json:"tanggal_check_out" gorm:"type:date"`
	Total_biaya       int                         `json:"total_biaya" gorm:"type:int" validate:"required"`
	Status_booking    statusbooking.StatusBooking `json:"status_booking"`
	Status_booking_id int                         `json:"status_booking_id" gorm:"foreignKey:status_booking_id" validate:"required"`
	Created_at        time.Time                   `json:"created_at" gorm:"autoCreateTime"`
	Updated_at        time.Time                   `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestCreateBooking struct {
	Kamar_id          int `json:"kamar_id" validate:"required"`
	User_id           int `json:"user_id"`
	Hotel_id          int `json:"hotel_id"`
	Jumlah_hari       int `json:"jumlah_hari" validate:"required,gt=0"`
	Total_biaya       int `json:"total_biaya"`
	Status_booking_id int `json:"status_booking_id"`
}

type RequestCreateBookingSwagger struct {
	Kamar_id    int `json:"kamar_id" validate:"required"`
	Jumlah_hari int `json:"user_id"`
}

type RequestAdmin struct {
	Booking_id        int    `json:"booking_id" validate:"required"`
	Kamar_id          int    `json:"kamar_id" validate:"required"`
	Tanggal_check_in  string `json:"tanggal_check_in"`
	Tanggal_check_out string `json:"tanggal_check_out"`
	Status_booking_id int    `json:"status_booking_id" validate:"required"`
	Status_kamar_id   int    `json:"status_kamar_id"`
}

type RequestUpdateBooking struct {
	Booking_id        int `json:"booking_id" validate:"required"`
	Kamar_id          int `json:"kamar_id" validate:"required"`
	Status_booking_id int `json:"status_booking_id" validate:"required"`
}

type InsertBooking struct {
	Booking_id        int    `json:"booking_id" validate:"required"`
	Kamar_id          int    `json:"kamar_id" validate:"required"`
	Tanggal_check_in  string `json:"tanggal_check_in"`
	Tanggal_check_out string `json:"tanggal_check_out"`
	Status_booking_id int    `json:"status_booking_id" validate:"required"`
}

type ResponseBookingUser struct {
	Booking_id        int                         `json:"booking_id" gorm:"primaryKey"`
	Kamar_id          int                         `json:"kamar_id" gorm:"foreignKey:kamar_id" validate:"required"`
	Kamar             kamar.Kamar                 `json:"kamar"`
	Hotel_id          int                         `json:"hotel_id" gorm:"foreignKey:hotel_id" validate:"required"`
	Hotel             hotel.Hotel                 `json:"hotel"`
	User_id           int                         `json:"user_id" gorm:"foreignKey:user_id" validate:"required"`
	Tanggal_check_in  string                      `json:"tanggal_check_in" gorm:"type:date" `
	Jumlah_hari       int                         `json:"jumlah_hari" gorm:"type:int"`
	Tanggal_check_out string                      `json:"tanggal_check_out" gorm:"type:date"`
	Total_biaya       int                         `json:"total_biaya" gorm:"type:int" validate:"required"`
	Status_booking    statusbooking.StatusBooking `json:"status_booking"`
	Status_booking_id int                         `json:"status_booking_id" gorm:"foreignKey:status_booking_id" validate:"required"`
	Created_at        time.Time                   `json:"created_at" gorm:"autoCreateTime"`
	Updated_at        time.Time                   `json:"updated_at" gorm:"autoUpdateTime"`
}

type ResponseBookingAdmin struct {
	Booking_id        int                         `json:"booking_id" gorm:"primaryKey"`
	Kamar_id          int                         `json:"kamar_id" gorm:"foreignKey:kamar_id" validate:"required"`
	Kamar             kamar.Kamar                 `json:"kamar"`
	Hotel_id          int                         `json:"hotel_id" gorm:"foreignKey:hotel_id" validate:"required"`
	Hotel             hotel.Hotel                 `json:"hotel"`
	User_id           int                         `json:"user_id" gorm:"foreignKey:user_id" validate:"required"`
	User              user.ResponseUser           `json:"user"`
	Tanggal_check_in  string                      `json:"tanggal_check_in" gorm:"type:date" `
	Jumlah_hari       int                         `json:"jumlah_hari" gorm:"type:int"`
	Tanggal_check_out string                      `json:"tanggal_check_out" gorm:"type:date"`
	Total_biaya       int                         `json:"total_biaya" gorm:"type:int" validate:"required"`
	Status_booking    statusbooking.StatusBooking `json:"status_booking"`
	Status_booking_id int                         `json:"status_booking_id" gorm:"foreignKey:status_booking_id" validate:"required"`
	Created_at        time.Time                   `json:"created_at" gorm:"autoCreateTime"`
	Updated_at        time.Time                   `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Booking) TableName() string {
	return "booking"
}
