package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppName string
	Port    string
	Debug   bool
	DB      DatabaseConfig
}

type DatabaseConfig struct {
	Name     string
	Username string
	Password string
	Host     string
}

func ReadConfiguration() (Configuration, error) {
	// 1. Menentukan nama file konfigurasi
	viper.SetConfigName(".env")

	// 2. Menentukan format file konfigurasi sebagai env
	viper.SetConfigType("env")

	// 3. Menentukan jalur direktori tempat file konfigurasi berada
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error membaca file konfigurasi: %s\n", err)
		return Configuration{}, err
	}
	return Configuration{
		AppName: viper.GetString("APP_NAME"),
		Port:    viper.GetString("PORT"),
		Debug:   viper.GetBool("DEBUG"),
		DB: DatabaseConfig{
			Name:     viper.GetString("DATABASE_NAME"),
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Host:     viper.GetString("DATABASE_HOST"),
		},
	}, nil
}
