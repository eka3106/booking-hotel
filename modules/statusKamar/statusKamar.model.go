package statuskamar

import "time"

type StatusKamar struct {
	Status_kamar_id   int       `json:"status_kamar_id" gorm:"primaryKey"`
	Status_kamar_nama string    `json:"status_kamar_nama" validate:"required,min=3"  gorm:"type:varchar(255)"`
	Created_at        time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
