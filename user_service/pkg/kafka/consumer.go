package kafka

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/user_service/config"
)

type MessageHandlerFunc func(ctx context.Context, topic string, message []byte) error

type Consumer interface {
	RegisterHandler(topic string, handlerFunc MessageHandlerFunc)
	Start(ctx context.Context) error
}

type consumer struct {
	saramaConsumer sarama.Consumer
	l              logger.Interface
	handlerFuncMap map[string]MessageHandlerFunc
}

func NewConsumer(cfg config.Config, l logger.Interface) (Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	saramaConsumer, err := sarama.NewConsumer([]string{""}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create sarama consumer: %w", err)
	}

	return &consumer{
		saramaConsumer: saramaConsumer,
		l:              l,
		handlerFuncMap: make(map[string]MessageHandlerFunc),
	}, nil
}

func (c consumer) RegisterHandler(topic string, handler MessageHandlerFunc) {
	c.handlerFuncMap[topic] = handler
}

func (c consumer) Start(ctx context.Context) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// for topic, handlerFunc := range c.handlerFuncMap {
	// 	go func() {
	// 		if err := c.saramaConsumer.C
	// 	} ()
	// }
	return nil
}
