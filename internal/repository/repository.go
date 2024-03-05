package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/model"
	"github.com/olegtemek/tg-store/internal/repository/user"
	"github.com/olegtemek/tg-store/internal/utils"
)

type User interface {
	Create(email string, password string) (*model.User, *utils.WrappError)
	GetByEmail(email string) (*model.User, *utils.WrappError)
}

type Repository struct {
	User
}

func New(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		User: user.NewRepository(log, db),
	}
}
