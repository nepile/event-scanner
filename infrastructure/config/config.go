package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error to read config file: %s", err)
	}

	AppConfig = Config{
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetString("DB_PORT"),
		DBUser: viper.GetString("DB_USER"),
		DBPass: viper.GetString("DB_PASS"),
		DBName: viper.GetString("DB_NAME"),
	}
}
