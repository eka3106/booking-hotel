package user

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"regexp"
	"strings"

	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

// CreateUser untuk membuat user godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param Register body RequestCreateUser true "Register"
// @Success 201 {string} string "Success Create User"
// @Router /1.0/register [post]
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
		query := `INSERT INTO users (nama, email_user, password_user) VALUES (?, ?, ?)`
		result := databases.DB.Exec(query, user.Nama, user.Email_user, user.Password_user)
		if result.Error != nil {
			return libs.ResponseError(c, result.Error.Error(), 500)
		} else {
			return libs.ResponseSuccess(c, "Success Create User", 201)
		}
	}
}

// CreateAdminHotel untuk membuat user admin hotel godoc
// @Summary Create Admin Hotel
// @Description Create Admin Hotel
// @Tags User
// @Accept json
// @Produce json
// @Param Register body RequestCreateAdminHotel true "Register"
// @Success 201 {string} string "Success Create User"
// @Router /1.0/register-admin-hotel [post]
func CreateAdminHotel(c *fiber.Ctx) error {
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
		query := `INSERT INTO users (nama, email_user, password_user, hotel_id, hak_akses_id) VALUES (?, ?, ?, ?, ?)`
		result := databases.DB.Exec(query, user.Nama, user.Email_user, user.Password_user, user.Hotel_id, 3)
		if result.Error != nil {
			return libs.ResponseError(c, result.Error.Error(), 500)
		} else {
			return libs.ResponseSuccess(c, "Success Create User", 201)
		}
	}
}

// Login untuk login user godoc
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param Login body RequestLogin true "Login"
// @Success 200 {string} string "Success Login"
// @Router /1.0/login [post]
func Login(c *fiber.Ctx) error {
	userReq := User{}
	if err := c.BodyParser(&userReq); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	email := libs.CleanText(userReq.Email_user)
	userDb := User{}
	query := `SELECT * FROM users WHERE email_user = ?`
	err := databases.DB.Raw(query, email).Scan(&userDb)
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
	query = `UPDATE users SET token =array_append(token,?)  WHERE user_id = ?`
	err = databases.DB.Exec(query, token, userDb.User_id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, token, 200)
}

// Logout untuk logout user godoc
// @Summary Logout User
// @Description Logout User
// @Tags User
// @Security Bearer
// @Produce json
// @Success 200 {string} string "Success Logout"
// @Router /1.0/logout [post]
func Logout(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	userClaims := claims.(*Claims)
	id := userClaims.User_id
	userDb := User{}
	err := databases.DB.Table("users").Where("user_id = ?", id).First(&userDb)
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "User Not Found", 404)
	}
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	token := libs.ExtractToken(c)
	queryRemoveToken := `UPDATE users SET token=array_remove(token, $1) WHERE user_id=$2`
	err = databases.DB.Exec(queryRemoveToken, token, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 500)
	}
	return libs.ResponseSuccess(c, "Success Logout", 200)
}

func createToken(userDb User) (string, error) {
	jwtClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":      userDb.User_id,
		"email":        userDb.Email_user,
		"name":         userDb.Nama,
		"hak_akses_id": userDb.Hak_akses_id,
		"hotel_id":     userDb.Hotel_id,
		"exp":          time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	token, err := jwtClaim.SignedString([]byte(viper.GetString("SECRET_JWT")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CheckToken(token string, id int) bool {
	var tokens string
	query := `SELECT token FROM users WHERE user_id = ?`
	err := databases.DB.Raw(query, id).Scan(&tokens).Error
	if err != nil {
		return false
	}
	re := regexp.MustCompile(`[{}]`)

	tokenList := strings.Split(re.ReplaceAllString(tokens, ""), ",")
	for _, t := range tokenList {
		if t == token {
			return true
		}
	}
	return false
}
