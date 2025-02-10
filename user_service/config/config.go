package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of application.
// The values read by viper from a config file or environment variables.
type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	DBHost               string        `mapstructure:"DB_HOST"`
	DBPort               int           `mapstructure:"DB_PORT"`
	DBUser               string        `mapstructure:"DB_USER"`
	DBPassword           string        `mapstructure:"DB_PASSWORD"`
	DBName               string        `mapstructure:"DB_NAME"`
	PublicKey            string        `mapstructure:"PUBLIC_KEY"`
	PrivateKey           string        `mapstructure:"PRIVATE_KEY"`
	RedisAddres          string        `mapstructure:"REDIS_ADDESS"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	KafkaBrokers         string        `mapstructure:"KAFKA_BROKERS"`
	KafkaClientId        string        `mapstructure:"KAFKA_CLIENT_ID"`
}

// LoadConfig reads configuration from file or enviroment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
