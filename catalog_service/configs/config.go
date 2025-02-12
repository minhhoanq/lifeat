package configs

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of application.
// The values read by viper from a config file or environment variables.
type Config struct {
	// GRPC     GRPC     `mapstructure:"grpc"`
	// Database Database `mapstructure:"database"`
	// Kafka    Kafka    `mapstructure:"kafka"`
	// Log      Log      `mapstructure:"log"`
	// Mail     Mail     `mapstructure:"mail"`
	Environment         string `mapstructure:"ENVIRONMENT"`
	GRPCServerAddress   string `mapstructure:"GRPC_SERVER_ADDRESS"`
	LogLevel            string `mapstructure:"LOG_LEVEL"`
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              int    `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	RedisAddres         string `mapstructure:"REDIS_ADDESS"`
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	KafkaBrokers        string `mapstructure:"KAFKA_BROKERS"`
	KafkaClientId       string `mapstructure:"KAFKA_CLIENT_ID"`
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
