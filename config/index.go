package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	// Load the configuration file
	loadConfig()
}

func loadConfig() {
	// Load the configuration file
	viper := viper.New()
	viper.SetConfigFile("../.env")
	if err:= viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
	}
	viper.AutomaticEnv()			
}