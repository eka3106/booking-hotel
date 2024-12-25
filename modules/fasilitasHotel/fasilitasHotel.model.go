package fasilitashotel

import "time"

type FasilitasHotel struct {
	Fasilitas_hotel_id int       `json:"fasilitas_hotel_id" gorm:"primaryKey"`
	Hotel_id           int       `json:"hotel_id" gorm:"foreignKey:Hotel_id" validate:"required"`
	Fasilitas_id       int       `json:"fasilitas_id" gorm:"foreignKey:Fasilitas_id" validate:"required"`
	Created_at         time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
