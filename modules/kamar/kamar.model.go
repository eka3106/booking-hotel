package kamar

import (
	"booking-hotel/modules/hotel"
	statuskamar "booking-hotel/modules/statusKamar"
	tipekamar "booking-hotel/modules/tipeKamar"
	"time"
)

type Kamar struct {
	Kamar_id        int       `json:"kamar_id" gorm:"primaryKey"`
	Hotel_id        int       `json:"hotel_id" gorm:"foreignKey:Hotel_id"`
	Nomor_kamar     string    `json:"nomor_kamar" gorm:"type:varchar(255)" validate:"required"`
	Tipe_kamar_id   int       `json:"tipe_kamar_id" gorm:"foreignKey:Tipe_kamar_id" validate:"required"`
	Harga           int       `json:"harga" gorm:"type:int" validate:"required,gt=50000"`
	Status_kamar_id int       `json:"status_kamar_id" gorm:"foreignKey:Status_kamar_id"`
	Created_at      time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type ResponseKamar struct {
	Kamar_id        int                     `json:"kamar_id" gorm:"primaryKey"`
	Hotel_id        int                     `json:"hotel_id" gorm:"foreignKey:Hotel_id"`
	Hotel           hotel.Hotel             `json:"hotel"`
	Nomor_kamar     string                  `json:"nomor_kamar" gorm:"type:varchar(255)" validate:"required"`
	Tipe_kamar_id   int                     `json:"tipe_kamar_id" gorm:"foreignKey:Tipe_kamar_id" validate:"required"`
	Tipe_kamar      tipekamar.TipeKamar     `json:"tipe_kamar"`
	Harga           int                     `json:"harga" gorm:"type:int" validate:"required,gt=50000"`
	Status_kamar    statuskamar.StatusKamar `json:"status_kamar"`
	Status_kamar_id int                     `json:"status_kamar_id" gorm:"foreignKey:Status_kamar_id"`
	Created_at      time.Time               `json:"created_at" gorm:"autoCreateTime"`
	Updated_at      time.Time               `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestKamar struct {
	Hotel_id        int    `json:"hotel_id" validate:"required"`
	Nomor_kamar     string `json:"nomor_kamar" validate:"required"`
	Tipe_kamar_id   int    `json:"tipe_kamar_id" validate:"required"`
	Harga           int    `json:"harga" validate:"required,gt=50000"`
	Status_kamar_id int    `json:"status_kamar_id" `
}

func (Kamar) TableName() string {
	return "kamar"
}
