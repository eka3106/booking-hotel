package pembayaran

import "time"

type Pembayaran struct {
	Pembayaran_id        int       `json:"pembayaran_id" gorm:"primaryKey"`
	Booking_id           int       `json:"booking_id" gorm:"foreignKey:Booking_id" validate:"required"`
	Total_pembayaran     int       `json:"total_pembayaran" gorm:"type:int" validate:"required"`
	Tanggal_pembayaran   time.Time `json:"tanggal_pembayaran" gorm:"autoCreateTime"`
	Status_pembayaran_id int       `json:"status_pembayaran_id" gorm:"foreignKey:Status_pembayaran_id" validate:"required"`
}
