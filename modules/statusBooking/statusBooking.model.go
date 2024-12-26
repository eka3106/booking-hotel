package statusbooking

import "time"

type StatusBooking struct {
	Status_booking_id   int       `json:"status_booking_id" gorm:"primaryKey"`
	Status_booking_nama string    `json:"status_booking_nama" validate:"required,min=3"  gorm:"type:varchar(255)"`
	Created_at          time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (StatusBooking) TableName() string {
	return "status_booking"
}

type RequestStatusBooking struct {
	Status_booking_nama string `json:"status_booking_nama" validate:"required"`
}
