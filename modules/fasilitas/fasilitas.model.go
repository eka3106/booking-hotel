package fasilitas

import "time"

type Fasilitas struct {
	Fasilitas_id    int       `json:"fasilitas_id" goorm:"primary_key"`
	Jenis_fasilitas string    `json:"jenis_fasilitas"  goorm:"type:varchar(255)" validate:"required"`
	Hotel_id        int       `json:"hotel_id" goorm:"type:int" validate:"required"`
	Deskripsi       string    `json:"deskripsi" goorm:"type:text" validate:"required, min=10"`
	Created_at      time.Time `json:"created_at" goorm:"type:timestamp"`
	Updated_at      time.Time `json:"updated_at" goorm:"type:timestamp"`
}
