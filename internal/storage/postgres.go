package storage

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/config"
)

func NewPostgresConnect(log *slog.Logger, cfg *config.Config) (*pgxpool.Pool, error) {

	for i := 0; i < 5; i++ {

		conn, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)

		if err != nil {
			log.Error("cannot connect to db", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if err := conn.Ping(context.Background()); err != nil {
			log.Error("cannot connect to db", err)
			time.Sleep(3 * time.Second)
			continue
		}

		return conn, nil
	}

	return nil, fmt.Errorf("cannot connect to db")
}
