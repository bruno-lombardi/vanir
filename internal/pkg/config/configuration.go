package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config *Configuration

type Configuration struct {
	Server      ServerConfiguration
	Database    DatabaseConfiguration
	HttpClients HttpClientConfiguration
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
	path := os.Getenv("ENV_FILE")
	if path != "" {
		configLogger.Printf("loading configuration from file %v\n", path)
		godotenv.Load(path)
	} else {
		configLogger.Printf("loading configuration from env\n")
		godotenv.Load()
	}

	configuration := &Configuration{
		Server: ServerConfiguration{
			Port:   os.Getenv("SERVER_PORT"),
			Secret: os.Getenv("SERVER_SECRET"),
		},
		Database: DatabaseConfiguration{
			Driver:       os.Getenv("DB_DRIVER"),
			DbName:       os.Getenv("DB_NAME"),
			Username:     os.Getenv("DB_USERNAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			MaxLifetime:  int(parseInt(os.Getenv("DB_MAX_LIFE_TIME"))),
			MaxOpenConns: int(parseInt(os.Getenv("DB_MAX_OPEN_CONNS"))),
			MaxIdleConns: int(parseInt(os.Getenv("DB_MAX_IDLE_CONNS"))),
		},
		HttpClients: HttpClientConfiguration{
			CryptoCompareBaseURL: os.Getenv("CRYPTO_COMPARE_BASE_URL"),
			CryptoCompareAPIKey:  os.Getenv("CRYPTO_COMPARE_API_KEY"),
		},
	}

	Config = configuration
}

func parseInt(value string) int64 {
	i, _ := strconv.ParseInt(value, 10, 0)
	return i
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
