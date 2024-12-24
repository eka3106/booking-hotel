package user

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if err := validate.Struct(user); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password_user), bcrypt.DefaultCost)
	if err != nil {
		return libs.ResponseError(c, err.Error(), 500)
	} else {
		user.Password_user = string(bytes)
		err := databases.DB.Create(&user)
		if err != nil {
			return libs.ResponseError(c, err.Error.Error(), 500)
		} else {
			return libs.ResponseSuccess(c, "Success Create User", 201)
		}
	}
}

func Login(c *fiber.Ctx) error {
	userReq := User{}
	if err := c.BodyParser(&userReq); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	email := libs.CleanText(userReq.Email_user)
	userDb := User{}
	err := databases.DB.Where("email_user = ?", email).First(&userDb)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Email or Password is Wrong", 400)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password_user), []byte(userReq.Password_user)); err != nil {
		return libs.ResponseError(c, "Email or Password is Wrong", 400)
	}
	token, errorJwt := createToken(userDb)
	if errorJwt != nil {
		return libs.ResponseError(c, errorJwt.Error(), 500)
	}
	userDb.Token = append(userDb.Token, token)
	err = databases.DB.Save(&userDb)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, userDb, 200)
}

func Logout(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	userClaims := claims.(Claims)
	id := userClaims.IdUser
	userDb := User{}
	err := databases.DB.Where("user_id = ?", id).First(&userDb)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "User Not Found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	token := libs.ExtractToken(c)
	for i, v := range userDb.Token {
		if v == token {
			userDb.Token = append(userDb.Token[:i], userDb.Token[i+1:]...)
			break
		}
	}
	err = databases.DB.Save(&userDb)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Success Logout", 200)
}

func createToken(userDb User) (string, error) {
	jwtClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"idUser": userDb.User_id,
		"email":  userDb.Email_user,
		"name":   userDb.Nama,
		"exp":    time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	token, err := jwtClaim.SignedString([]byte(viper.GetString("SECRET_JWT")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CheckToken(token string, id int) bool {
	userDb := User{}
	err := databases.DB.Where("user_id = ?", id).First(&userDb)
	if err.RowsAffected == 0 {
		return false
	}
	if err.Error != nil {
		return false
	}
	for _, v := range userDb.Token {
		if v == token {
			return true
		}
	}
	return false
}
