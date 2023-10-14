package app

import (
	"log"
	"url-short/internal/config"
	"url-short/internal/logging"
)

func Run() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error occured while parsing configs: %v", err)
	}

	_ = config.GetConfig()

	_ = logging.InitLogger()

	// TODO: init database

	// TODO: init layers

	// TODO: init server
}
