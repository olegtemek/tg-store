package user

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/model"
	"github.com/olegtemek/tg-store/internal/utils"
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

func (r *Repository) Create(email string, password string) (*model.User, *utils.WrappError) {
	user := &model.User{}
	q := `INSERT INTO Users (email, password) VALUES ($1, $2) RETURNING id, email, password;`
	err := r.db.QueryRow(context.Background(), q, email, password).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return user, &utils.WrappError{Err: err}
	}
	return user, nil
}

func (r *Repository) GetByEmail(email string) (*model.User, *utils.WrappError) {
	user := &model.User{}
	q := `SELECT id, email, password FROM Users WHERE email = $1`

	err := r.db.QueryRow(context.Background(), q, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return user, &utils.WrappError{Err: err}
	}
	return user, nil

}
