package configs

type Kafka struct {
	Brokers  []string `mapstructure:"KAFKA_BROKERS"`
	ClientID string   `mapstructure:"KAFKA_CLIENT_ID"`
}
