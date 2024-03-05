package user

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/service"
	"github.com/olegtemek/tg-store/internal/utils"
	pb "github.com/olegtemek/tg-store/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	log                *slog.Logger
	accessTokenSecret  string
	refreshTokenSecret string
	service            service.User
	pb.UnimplementedUserServiceServer
}

func NewGRPCHandler(log *slog.Logger, service *service.User, accessTokenSecret string, refreshTokenSecret string) *Handler {
	return &Handler{
		log:                log,
		accessTokenSecret:  accessTokenSecret,
		refreshTokenSecret: refreshTokenSecret,
		service:            *service,
	}
}

func (h *Handler) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	dto := &dto.UserRegistration{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	user, wrapErr := h.service.Registration(dto)
	if wrapErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", wrapErr.Error()))
	}

	accessToken, err := utils.GenerateAccessToken(h.accessTokenSecret, int(user.Id))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	refreshToken, err := utils.GenerateRefreshToken(h.refreshTokenSecret, int(user.Id))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	return &pb.RegistrationResponse{
		User: &pb.UserModel{
			Id:    int64(user.Id),
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	dto := &dto.UserLogin{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	user, wrapErr := h.service.Login(dto)
	if wrapErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", wrapErr.Error()))
	}

	accessToken, err := utils.GenerateAccessToken(h.accessTokenSecret, int(user.Id))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	refreshToken, err := utils.GenerateRefreshToken(h.refreshTokenSecret, int(user.Id))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	return &pb.LoginResponse{
		User: &pb.UserModel{
			Id:    int64(user.Id),
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h *Handler) GetProfile(ctx context.Context, req *pb.Empty) (*pb.ProfileResponse, error) {
	user := ctx.Value("user").(*pb.UserModel)

	return &pb.ProfileResponse{
		User: &pb.UserModel{
			Id:    int64(user.Id),
			Email: user.Email,
		},
	}, nil
}
