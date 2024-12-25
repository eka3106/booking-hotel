package kamar

type Kamar struct {
	Kamar_id        int    `json:"kamar_id" gorm:"primaryKey"`
	Hotel_id        int    `json:"hotel_id" gorm:"foreignKey:Hotel_id"`
	Nomor_kamar     string `json:"nomor_kamar" gorm:"type:varchar(255)" validate:"required"`
	Tipe_kamar_id   string `json:"tipe_kamar_id" gorm:"foreignKey:Tipe_kamar_id" validate:"required"`
	Harga           int    `json:"harga" gorm:"type:int" validate:"required, gt=50000"`
	Status_kamar_id string `json:"status_kamar_id" gorm:"foreignKey:Status_kamar_id" validate:"required"`
	Created_at      string `json:"created_at" gorm:"autoCreateTime"`
	Updated_at      string `json:"updated_at" gorm:"autoUpdateTime"`
}
