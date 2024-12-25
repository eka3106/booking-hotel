package booking

import "time"

type Booking struct {
	Booking_id        int       `json:"booking_id" gorm:"primaryKey"`
	Kamar_id          int       `json:"kamar_id" gorm:"foreignKey:kamar_id" validate:"required"`
	Hotel_id          int       `json:"hotel_id" gorm:"foreignKey:hotel_id" validate:"required"`
	User_id           int       `json:"user_id" gorm:"foreignKey:user_id" validate:"required"`
	Tanggal_check_in  string    `json:"tanggal_check_in" gorm:"type:date" `
	Tanggal_check_out string    `json:"tanggal_check_out" gorm:"type:date"`
	Total_biaya       int       `json:"total_biaya" gorm:"type:int" validate:"required"`
	Status_booking_id int       `json:"status_booking_id" gorm:"foreignKey:status_booking_id" validate:"required"`
	Created_at        time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestCheckIn struct {
	Booking_id       int    `json:"booking_id" validate:"required"`
	Tanggal_check_in string `json:"tanggal_check_in" validate:"required,date"`
}

type RequestCheckOut struct {
	Booking_id        int    `json:"booking_id" validate:"required"`
	Tanggal_check_out string `json:"tanggal_check_out" validate:"required,date"`
}
