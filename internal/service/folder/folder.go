package folder

import (
	"log/slog"

	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/repository"
	"github.com/olegtemek/tg-store/internal/utils"
	pb "github.com/olegtemek/tg-store/proto"
)

type Service struct {
	log  *slog.Logger
	repo repository.Folder
}

func NewService(log *slog.Logger, repo *repository.Folder) *Service {
	return &Service{
		log:  log,
		repo: *repo,
	}
}

func (s *Service) Create(dto *dto.FolderCreate) (*pb.FolderModel, *utils.WrappError) {
	folder, err := s.repo.Create(dto.Title, dto.UserId)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (s *Service) Update(dto *dto.FolderUpdate) (*pb.FolderModel, *utils.WrappError) {
	folder, err := s.repo.Update(dto.Id, dto.Title, dto.UserId)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (s *Service) GetAll(dto *dto.FolderGetAll) ([]*pb.FolderModel, *utils.WrappError) {
	folders, err := s.repo.GetAll(dto.UserId)
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func (s *Service) Delete(dto *dto.FolderDelete) (*pb.FolderModel, *utils.WrappError) {
	folder, err := s.repo.Delete(dto.Id, dto.UserId)
	if err != nil {
		return nil, err
	}
	return folder, nil
}
