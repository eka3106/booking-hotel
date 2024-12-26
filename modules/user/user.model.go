package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	User_id       int       `json:"user_id" gorm:"primaryKey;autoIncrement;not null"`
	Email_user    string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password_user string    `json:"password" gorm:"not null" validate:"required,min=8"`
	Nama          string    `json:"nama" gorm:"not null" validate:"required,min=3"`
	Hak_akses_id  int       `json:"hak_akses_id" gorm:"not null,default:2"`
	Token         string    `json:"token"`
	Hotel_id      int       `json:"hotel_id"`
	Created_at    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Claims struct {
	User_id      int    `json:"user_id"`
	Name         string `json:"name"`
	Hak_akses_id int    `json:"hak_akses_id"`
	Hotel_id     int    `json:"hotel_id"`
	Email        string `json:"email"`
	jwt.RegisteredClaims
}
