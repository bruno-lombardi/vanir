package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration   `mapstructure:",squash"`
	Database DatabaseConfiguration `mapstructure:",squash"`
}

type DatabaseConfiguration struct {
	Driver       string `mapstructure:"DB_DRIVER"`
	DbName       string `mapstructure:"DB_NAME"`
	Username     string `mapstructure:"DB_USERNAME"`
	Password     string `mapstructure:"DB_PASSWORD"`
	Host         string `mapstructure:"DB_HOST"`
	Port         string `mapstructure:"DB_PORT"`
	MaxLifetime  int    `mapstructure:"DB_MAX_LIFE_TIME"`
	MaxOpenConns int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns int    `mapstructure:"DB_MAX_IDLE_CONNS"`
}

type ServerConfiguration struct {
	Port   string `mapstructure:"SERVER_PORT"`
	Secret string `mapstructure:"SERVER_SECRET"`
}

// SetupDB initialize configuration
func Setup() {
	var configuration *Configuration

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
	fmt.Printf("config: %v", configuration)
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
