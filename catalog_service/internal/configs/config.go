package configs

import (
	"fmt"
	"os"

	"github.com/minhhoanq/lifeat/order_service/configs"
	"gopkg.in/yaml.v2"
)

type ConfigFilePath = string

type Config struct {
	GRPC     GRPC     `yaml:"grpc"`
	Database Database `yaml:"database"`
	Kafka    Kafka    `yaml:"kafka"`
	Log      Log      `yaml:"log"`
	Mail     Mail     `yaml:"mail"`
}

func NewConfig(configFilePath ConfigFilePath) (Config, error) {
	var (
		configBytes []byte = configs.DefaultConfigBytes
		config      Config
		err         error
	)

	if configFilePath != "" {
		configBytes, err = os.ReadFile(string(configFilePath))
		if err != nil {
			return Config{}, fmt.Errorf("error reading configuration file: %w", err)
		}
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshal configuration file %w", err)
	}

	return config, nil
}
