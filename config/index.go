package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {

	// viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
	}

	viper.SetDefault("PORT", os.Getenv("PORT"))
	viper.SetDefault("DB_HOST", os.Getenv("DB_HOST"))
	viper.SetDefault("DB_PORT", os.Getenv("DB_PORT"))
	viper.SetDefault("DB_USER", os.Getenv("DB_USER"))
	viper.SetDefault("DB_PASSWORD", os.Getenv("DB_PASSWORD"))
	viper.SetDefault("DB_NAME", os.Getenv("DB_NAME"))
	viper.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))
}
