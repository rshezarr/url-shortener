package app

import (
	"log"
	"url-short/internal/config"
)

func Run() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error occured while parsing configs: %v", err)
	}

	_ = config.GetConfig()

	// TODO: init logger

	// TODO: init database

	// TODO: init layers

	// TODO: init server
}
