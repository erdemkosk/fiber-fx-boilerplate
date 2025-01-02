package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig
	Mongo  MongoConfig
	Logger LoggerConfig
}

type AppConfig struct {
	Env            string
	Port           int
	SwaggerEnabled bool
	RateLimit      int
	RateWindow     time.Duration
}

type MongoConfig struct {
	URL          string
	QueryTimeout time.Duration
}

type LoggerConfig struct {
	Level string
}

func LoadEnv() *Config {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading .env file:", err)
	}

	config := &Config{
		App: AppConfig{
			Env:            viper.GetString("APP_ENV"),
			Port:           viper.GetInt("PORT"),
			SwaggerEnabled: viper.GetBool("SWAGGER_ENABLED"),
			RateLimit:      viper.GetInt("RATE_LIMIT"),
			RateWindow:     viper.GetDuration("RATE_WINDOW"),
		},
		Mongo: MongoConfig{
			URL:          viper.GetString("MONGO_DB_URL"),
			QueryTimeout: viper.GetDuration("QUERY_TIMEOUT"),
		},
		Logger: LoggerConfig{
			Level: viper.GetString("LOG_LEVEL"),
		},
	}

	return config
}
