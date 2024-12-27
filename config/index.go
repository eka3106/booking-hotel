package config

import (
	"github.com/spf13/viper"
)

func init() {

	// viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("Error reading config file", err)
	// }
}
