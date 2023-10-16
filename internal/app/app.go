package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"url-short/internal"
	"url-short/internal/config"
	v1 "url-short/internal/http/v1"
	"url-short/internal/repository"
	"url-short/internal/service"
	"url-short/pkg/database/postgre"
)

func Run() {
	cfg := config.GetConfig()
	cfg.Logger.Info("config initialized")

	psql, err := postgre.ConnectPSQL(cfg)
	if err != nil {
		cfg.Logger.Info("database connection failed: ", slog.String("error", err.Error()))
		return
	}
	cfg.Logger.Info("database connected successfully")

	repo := repository.NewRepository(psql.DB)
	svc := service.NewService(repo)
	ctrl := v1.NewController(svc)

	srv := internal.NewServer(cfg, ctrl.InitRoutes())
	cfg.Logger.Info("server created successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		cfg.Logger.Info("server started", slog.String("server-port", cfg.HTTP.Port))
		srv.Run()
	}()

	select {
	case sig := <-quit:
		cfg.Logger.Info("app: signal accepted", slog.String("signal", sig.String()))
	case err := <-srv.ServerErrorNotify:
		cfg.Logger.Info("app: signal accepted", slog.String("error", err.Error()))
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		cfg.Logger.Error("error while shutting down server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown")

	if err := psql.CloseBD(); err != nil {
		cfg.Logger.Error("error while closing database", slog.String("error", err.Error()))
	}
	slog.Info("db closing")
}
