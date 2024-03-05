package user

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/model"
)

type Repository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		log: log,
		db:  db,
	}
}

func (r *Repository) Create(email string, login string, password string) (*model.User, error) {
	panic("IMPLEMENT ME")
	return nil, nil
}
