package user

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/model"
	"github.com/olegtemek/tg-store/internal/service"
	"github.com/olegtemek/tg-store/internal/utils"
	tgstorev1 "github.com/olegtemek/tg-store/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	log                *slog.Logger
	accessTokenSecret  string
	refreshTokenSecret string
	service            service.User
	tgstorev1.UnimplementedUserServiceServer
}

func NewGRPCHandler(log *slog.Logger, service *service.User, accessTokenSecret string, refreshTokenSecret string) *Handler {
	return &Handler{
		log:                log,
		accessTokenSecret:  accessTokenSecret,
		refreshTokenSecret: refreshTokenSecret,
		service:            *service,
	}
}

func (h *Handler) Registration(ctx context.Context, req *tgstorev1.RegistrationRequest) (*tgstorev1.RegistrationResponse, error) {
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

	accessToken, err := utils.GenerateAccessToken(h.accessTokenSecret, user.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	refreshToken, err := utils.GenerateRefreshToken(h.refreshTokenSecret, user.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	return &tgstorev1.RegistrationResponse{
		User: &tgstorev1.User{
			Id:    int64(user.Id),
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h *Handler) Login(ctx context.Context, req *tgstorev1.LoginRequest) (*tgstorev1.LoginResponse, error) {
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

	accessToken, err := utils.GenerateAccessToken(h.accessTokenSecret, user.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	refreshToken, err := utils.GenerateRefreshToken(h.refreshTokenSecret, user.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	return &tgstorev1.LoginResponse{
		User: &tgstorev1.User{
			Id:    int64(user.Id),
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h *Handler) GetProfile(ctx context.Context, req *tgstorev1.Empty) (*tgstorev1.ProfileResponse, error) {
	user := ctx.Value("user").(*model.User)

	return &tgstorev1.ProfileResponse{
		User: &tgstorev1.User{
			Id:    int64(user.Id),
			Email: user.Email,
		},
	}, nil
}
