package user

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/utils"
	pb "github.com/olegtemek/tg-store/proto"
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

func (r *Repository) Create(email string, password string) (*pb.UserModel, *utils.WrappError) {
	user := &pb.UserModel{}
	q := `INSERT INTO Users (email, password) VALUES ($1, $2) RETURNING id, email, password;`
	err := r.db.QueryRow(context.Background(), q, email, password).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		r.log.Error("ERROR",
			slog.String("SQL", q),
			slog.Any("ERROR", err),
		)
		return nil, &utils.WrappError{Err: err}
	}
	return user, nil
}

func (r *Repository) GetByEmail(email string) (*pb.UserModel, *utils.WrappError) {
	user := &pb.UserModel{}
	q := `SELECT id, email, password FROM Users WHERE email = $1`

	err := r.db.QueryRow(context.Background(), q, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		r.log.Error("ERROR",
			slog.String("SQL", q),
			slog.Any("ERROR", err),
		)
		return nil, &utils.WrappError{Err: err}
	}
	return user, nil
}

func (r *Repository) GetById(id int) (*pb.UserModel, *utils.WrappError) {
	user := &pb.UserModel{}
	q := `SELECT id, email, password FROM Users WHERE id = $1`

	err := r.db.QueryRow(context.Background(), q, id).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		r.log.Error("ERROR",
			slog.String("SQL", q),
			slog.Any("ERROR", err),
		)
		return nil, &utils.WrappError{Err: err}
	}
	return user, nil

}
