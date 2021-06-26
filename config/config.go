package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {

	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./world")
	// viper.SetConfigName("config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

}
