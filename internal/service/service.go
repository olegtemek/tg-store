package service

import (
	"log/slog"

	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/repository"
	"github.com/olegtemek/tg-store/internal/service/folder"
	"github.com/olegtemek/tg-store/internal/service/user"
	"github.com/olegtemek/tg-store/internal/utils"
	pb "github.com/olegtemek/tg-store/proto"
)

type User interface {
	Login(dto *dto.UserLogin) (*pb.UserModel, *utils.WrappError)
	Registration(dto *dto.UserRegistration) (*pb.UserModel, *utils.WrappError)
	GetById(userId int) (*pb.UserModel, *utils.WrappError)
}

type Folder interface {
	Create(dto *dto.FolderCreate) (*pb.FolderModel, *utils.WrappError)
	Update(dto *dto.FolderUpdate) (*pb.FolderModel, *utils.WrappError)
	GetAll(dto *dto.FolderGetAll) ([]*pb.FolderModel, *utils.WrappError)
	Delete(dto *dto.FolderDelete) (*pb.FolderModel, *utils.WrappError)
}

type Service struct {
	log *slog.Logger
	User
	Folder
}

func New(log *slog.Logger, repos *repository.Repository) *Service {
	return &Service{
		log:    log,
		User:   user.NewSevice(log, &repos.User),
		Folder: folder.NewService(log, &repos.Folder),
	}
}
