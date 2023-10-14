package app

import (
	"fmt"
	"log"
	"url-short/internal/config"
	"url-short/internal/logging"
	"url-short/pkg/database/postgre"
)

func Run() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error occured while parsing configs: %v", err)
	}

	cfg := config.GetConfig()

	logger := logging.InitLogger()

	_, err := postgre.ConnectDB(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("database connection failed: %v", err))
		return
	}
	logger.Info("database connected successful")

	// TODO: init layers

	// TODO: init server
}
