package main

import (
	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/user_service/config"
	"github.com/minhhoanq/lifeat/user_service/internal/app/rest"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// init logger
	logger.Setup(cfg.Environment, cfg.LogLevel)

	// start rest server
	rest.RunRestServer(cfg)

}
