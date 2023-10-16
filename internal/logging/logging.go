package logging

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("logger initialized")

	return logger
}
