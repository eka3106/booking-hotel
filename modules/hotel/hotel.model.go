package hotel

import "time"

type Hotel struct {
	Hotel_id     int       `json:"hotel_id" gorm:"primaryKey"`
	Nama_hotel   string    `json:"nama_hotel" gorm:"type:varchar(255)" validate:"required,min=3"`
	Alamat_hotel string    `json:"alamat_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Telp_hotel   string    `json:"telp_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Email_hotel  string    `json:"email_hotel" gorm:"type:varchar(255)" validate:"required,email"`
	Created_at   time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ResponseHotel struct {
	Hotel_id     int    `json:"hotel_id" gorm:"primaryKey"`
	Nama_hotel   string `json:"nama_hotel" gorm:"type:varchar(255)" validate:"required,min=3"`
	Alamat_hotel string `json:"alamat_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Rating_hotel string `json:"rating_hotel" gorm:"type:varchar(255)" validate:"required,min=1"`
	Telp_hotel   string `json:"telp_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Email_hotel  string `json:"email_hotel" gorm:"type:varchar(255)" validate:"required,email"`
	Created_at   string `json:"created_at" gorm:"autoCreateTime"`
	Updated_at   string `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Hotel) TableName() string {
	return "hotel"
}

type RequestHotel struct {
	Nama_hotel   string `json:"nama_hotel" gorm:"type:varchar(255)" validate:"required,min=3"`
	Alamat_hotel string `json:"alamat_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Telp_hotel   string `json:"telp_hotel" gorm:"type:varchar(255)" validate:"required,min=10"`
	Email_hotel  string `json:"email_hotel" gorm:"type:varchar(255)" validate:"required,email"`
}