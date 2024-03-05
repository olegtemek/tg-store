package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/model"
	"github.com/olegtemek/tg-store/internal/repository/user"
)

type User interface {
	Create(email string, login string, password string) (*model.User, error)
}

type Repository struct {
	User
}

func New(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		User: user.NewRepository(log, db),
	}
}
