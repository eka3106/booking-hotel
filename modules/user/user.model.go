package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	User_id       int       `json:"user_id" gorm:"primaryKey;autoIncrement;not null"`
	Email_user    string    `json:"email_user" gorm:"unique;not null" validate:"required,email"`
	Password_user string    `json:"password_user" gorm:"not null" validate:"required,min=8"`
	Nama          string    `json:"name_user" gorm:"not null" validate:"required,min=3"`
	Hak_akses_id  int       `json:"hak_akses_id" gorm:"not null, default:2"`
	Token         []string  `json:"token" gorm:""`
	Created_at    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Claims struct {
	IdUser int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
