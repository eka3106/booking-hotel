package databases

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	SqlDb *sql.DB
)

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", viper.GetString("DB_HOST"), viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_NAME"), viper.GetInt("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error connecting to database", err)
	}
	fmt.Println("Database connected")
	DB = db
	SqlDb = sqlDB
}
