package pembayaran

import (
	"booking-hotel/modules/booking"
	"time"
)

type Pembayaran struct {
	Pembayaran_id      int       `json:"pembayaran_id" gorm:"primaryKey"`
	Booking_id         int       `json:"booking_id" gorm:"foreignKey:Booking_id" validate:"required"`
	Total_pembayaran   int       `json:"total_pembayaran" gorm:"type:int" validate:"required"`
	Tanggal_pembayaran time.Time `json:"tanggal_pembayaran" gorm:"autoCreateTime"`
}

type ResponsePembayaran struct {
	Pembayaran_id      int             `json:"pembayaran_id" gorm:"primaryKey"`
	Booking_id         int             `json:"booking_id" gorm:"foreignKey:Booking_id" validate:"required"`
	Booking            booking.Booking `json:"booking"`
	Total_pembayaran   int             `json:"total_pembayaran" gorm:"type:int" validate:"required"`
	Tanggal_pembayaran time.Time       `json:"tanggal_pembayaran" gorm:"autoCreateTime"`
}
