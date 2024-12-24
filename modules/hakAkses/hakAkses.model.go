package hakakses

import "time"

type HakAkses struct {
	Hak_akses_id int       `json:"hak_akses_id" gorm:"primary_key;AUTO_INCREMENT"`
	Hak_akses    string    `json:"hak_akses" validate:"required,alphanum,min=3,max=50"`
	Created_at   time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
