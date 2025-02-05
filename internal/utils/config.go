package utils

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from a YAML file.
func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}
