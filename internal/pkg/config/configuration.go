package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server      ServerConfiguration     `mapstructure:",squash"`
	Database    DatabaseConfiguration   `mapstructure:",squash"`
	HttpClients HttpClientConfiguration `mapstructure:",squash"`
}

type HttpClientConfiguration struct {
	CryptoCompareBaseURL string `mapstructure:"CRYPTO_COMPARE_BASE_URL"`
	CryptoCompareAPIKey  string `mapstructure:"CRYPTO_COMPARE_API_KEY"`
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

var configLogger *log.Logger

// SetupDB initialize configuration
func Setup() {
	configLogger := log.Default()
	configLogger.SetPrefix("[CONFIG]: ")

	var configuration *Configuration
	path := os.Getenv("ENV_FILE")
	if path != "" {
		configLogger.Printf("loading env from file: %v\n", path)
	}

	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		configLogger.Printf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
