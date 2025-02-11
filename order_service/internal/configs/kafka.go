package configs

type Kafka struct {
	Brokers  []string `yaml:"brokers"`
	ClientID string   `yaml:"client_id"`
}
