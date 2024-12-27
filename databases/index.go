package databases

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	SqlDb *sql.DB
)

func init() {
	viper.SetConfigFile(".env")
    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    // Membuat string koneksi
    serviceURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require",
        viper.GetString("DB_USER"),
        viper.GetString("DB_PASSWORD"),
        viper.GetString("DB_HOST"),
        viper.GetInt("DB_PORT"),
        viper.GetString("DB_NAME"),
    )

    conn, err := url.Parse(serviceURI)
    if err != nil {
        log.Fatalf("Error parsing service URI, %s", err)
    }
    conn.RawQuery = "sslmode=verify-ca&sslrootcert=ca.pem"

    // Membuka koneksi dengan GORM
    db, err := gorm.Open(postgres.Open(conn.String()), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database, %s", err)
    }

    // Mendapatkan koneksi SQL dari GORM
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Error getting SQL DB from GORM, %s", err)
    }

	DB = db
	SqlDb = sqlDB
}
