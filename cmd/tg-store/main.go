package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/olegtemek/tg-store/internal/config"
	"github.com/olegtemek/tg-store/internal/repository"
	"github.com/olegtemek/tg-store/internal/service"
	"github.com/olegtemek/tg-store/internal/storage"
	appgrpc "github.com/olegtemek/tg-store/internal/transport/grpc"
)

func main() {

	cfg := config.New()

	log := setUpLogger(cfg.Env)
	db, err := storage.NewPostgresConnect(log, cfg)
	if err != nil {
		panic("Cannot connect to database")
	}

	repositories := repository.New(log, db)

	services := service.New(log, repositories)

	application := appgrpc.New(log, cfg, services)

	go func() {
		application.ListenAndServe()
	}()
	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
	log.Info("Gracefully stopped")
}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
