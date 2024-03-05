package user

import (
	"log/slog"

	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/model"
	"github.com/olegtemek/tg-store/internal/repository"
)

type Service struct {
	log  *slog.Logger
	repo *repository.User
}

func NewSevice(log *slog.Logger, repo *repository.User) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) Login() {
	panic("Implement me")
}
func (s *Service) Registration(dto *dto.UserRegistration) (*model.User, error) {
	panic("Implement me")
}
