package config

type Config struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	Port       int    `mapstructure:"PORT"`
	MongoDbUrl string `mapstructure:"MONGO_DB_URL"`
}
