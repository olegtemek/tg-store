package user

import (
	"log/slog"

	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/repository"
	"github.com/olegtemek/tg-store/internal/utils"
	pb "github.com/olegtemek/tg-store/proto"
)

type Service struct {
	log  *slog.Logger
	repo repository.User
}

func NewSevice(log *slog.Logger, repo *repository.User) *Service {
	return &Service{
		log:  log,
		repo: *repo,
	}
}

func (s *Service) Login(dto *dto.UserLogin) (*pb.UserModel, *utils.WrappError) {
	user, wrappErr := s.repo.GetByEmail(dto.Email)
	if wrappErr != nil {
		return nil, wrappErr
	}

	matchPass, err := utils.ComparePassword(user.Password, dto.Password)
	if err != nil || !matchPass {
		return nil, &utils.WrappError{Err: err}
	}

	return user, nil
}

func (s *Service) Registration(dto *dto.UserRegistration) (*pb.UserModel, *utils.WrappError) {

	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return nil, &utils.WrappError{Err: err}
	}

	user, wrapErr := s.repo.Create(dto.Email, hashedPassword)
	if wrapErr != nil {
		return nil, wrapErr
	}

	return user, nil
}

func (s *Service) GetById(userId int) (*pb.UserModel, *utils.WrappError) {

	user, wrapErr := s.repo.GetById(userId)
	if wrapErr != nil {
		return nil, wrapErr
	}

	return user, nil
}
