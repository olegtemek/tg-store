package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/repository/folder"
	"github.com/olegtemek/tg-store/internal/repository/user"
	"github.com/olegtemek/tg-store/internal/utils"
	pb "github.com/olegtemek/tg-store/proto"
)

type User interface {
	Create(email string, password string) (*pb.UserModel, *utils.WrappError)
	GetByEmail(email string) (*pb.UserModel, *utils.WrappError)
	GetById(id int) (*pb.UserModel, *utils.WrappError)
}

type Folder interface {
	Create(title string, userId int) (*pb.FolderModel, *utils.WrappError)
	Update(id int, title string, userId int) (*pb.FolderModel, *utils.WrappError)
	GetAll(userId int) ([]*pb.FolderModel, *utils.WrappError)
	Delete(id int, userId int) (*pb.FolderModel, *utils.WrappError)
}

type Repository struct {
	User
	Folder
}

func New(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		User:   user.NewRepository(log, db),
		Folder: folder.NewRepository(log, db),
	}
}
