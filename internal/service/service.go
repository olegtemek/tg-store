package service

import (
	"log/slog"

	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/model"
	"github.com/olegtemek/tg-store/internal/repository"
	"github.com/olegtemek/tg-store/internal/service/user"
	"github.com/olegtemek/tg-store/internal/utils"
)

type User interface {
	Login(dto *dto.UserLogin) (*model.User, *utils.WrappError)
	Registration(dto *dto.UserRegistration) (*model.User, *utils.WrappError)
	GetById(userId int) (*model.User, *utils.WrappError)
}
type Service struct {
	log *slog.Logger
	User
}

func New(log *slog.Logger, repos *repository.Repository) *Service {
	return &Service{
		log:  log,
		User: user.NewSevice(log, &repos.User),
	}
}
