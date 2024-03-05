package service

import (
	"log/slog"

	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/model"
	"github.com/olegtemek/tg-store/internal/repository"
	"github.com/olegtemek/tg-store/internal/service/user"
)

type User interface {
	Login()
	Registration(dto *dto.UserRegistration) (*model.User, error)
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
