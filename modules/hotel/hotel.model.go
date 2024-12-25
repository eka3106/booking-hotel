package hotel

type Hotel struct {
	Hotel_id     int    `json:"hotel_id" gorm:"primaryKey"`
	Nama_hotel   string `json:"nama_hotel" gorm:"type:varchar(255)" validate:"required,min=3"`
	Alamat_hotel string `json:"alamat_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Telp_hotel   string `json:"telp_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Email_hotel  string `json:"email_hotel" gorm:"type:varchar(255)" validate:"required,email"`
	Created_at   string `json:"created_at" gorm:"autoCreateTime"`
	Updated_at   string `json:"updated_at" gorm:"autoUpdateTime"`
}
