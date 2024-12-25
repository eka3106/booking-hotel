package tipekamar

import "time"

type TipeKamar struct {
	Tipe_kamar_id   int       `json:"tipe_kamar_id" gorm:"primaryKey"`
	Tipe_kamar_nama string    `validate:"required,min=3" json:"tipe_kamar_nama"`
	Created_at      time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
