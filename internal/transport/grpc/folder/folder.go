package folder

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/service"
	pb "github.com/olegtemek/tg-store/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	log     *slog.Logger
	service service.Folder
	pb.UnimplementedFolderServiceServer
}

func NewGRPCHandler(log *slog.Logger, service *service.Folder) *Handler {
	return &Handler{
		log:     log,
		service: *service,
	}
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateFolderRequest) (*pb.CreateFolderResponse, error) {
	user := ctx.Value("user").(*pb.UserModel)
	dto := &dto.FolderCreate{
		Title:  req.GetTitle(),
		UserId: int(user.Id),
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	folder, wrappErr := h.service.Create(dto)
	if wrappErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", wrappErr.Error()))
	}

	return &pb.CreateFolderResponse{
		Folder: &pb.FolderModel{
			Id:    int64(folder.Id),
			Title: folder.Title,
		},
	}, nil
}

func (h *Handler) Update(ctx context.Context, req *pb.UpdateFolderRequest) (*pb.UpdateFolderResponse, error) {
	user := ctx.Value("user").(*pb.UserModel)
	dto := &dto.FolderUpdate{
		Id:     int(req.Id),
		Title:  req.GetTitle(),
		UserId: int(user.Id),
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	folder, wrappErr := h.service.Update(dto)
	if wrappErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", wrappErr.Error()))
	}

	return &pb.UpdateFolderResponse{
		Folder: folder,
	}, nil
}

func (h *Handler) GetAll(ctx context.Context, req *pb.Empty) (*pb.GetAllFolderResponse, error) {
	user := ctx.Value("user").(*pb.UserModel)
	dto := &dto.FolderGetAll{
		UserId: int(user.Id),
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}
	folders, wrappErr := h.service.GetAll(dto)
	if wrappErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", wrappErr.Error()))
	}

	return &pb.GetAllFolderResponse{
		Folders: folders,
	}, nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteFolderRequest) (*pb.DeleteFolderResponse, error) {
	user := ctx.Value("user").(*pb.UserModel)
	dto := &dto.FolderDelete{
		Id:     int(req.Id),
		UserId: int(user.Id),
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}
	folder, wrappErr := h.service.Delete(dto)
	if wrappErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", wrappErr.Error()))
	}

	return &pb.DeleteFolderResponse{
		Folder: folder,
	}, nil
}
