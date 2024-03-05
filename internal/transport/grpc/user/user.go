package user

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/olegtemek/tg-store/internal/dto"
	"github.com/olegtemek/tg-store/internal/service"
	tgstorev1 "github.com/olegtemek/tg-store/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	tgstorev1.UnimplementedUserServer
	log     *slog.Logger
	service *service.User
}

func NewGRPCHandler(log *slog.Logger, service *service.User) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) Registration(ctx context.Context, req *tgstorev1.RegistrationRequest) (*tgstorev1.RegistrationResponse, error) {
	dto := &dto.UserRegistration{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Login:    req.GetLogin(),
	}
	if err := validator.New().Struct(dto); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("error: %s", err.Error()))
	}

	user, err := (*h.service).Registration(dto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error: %s", err.Error()))
	}

	return &tgstorev1.RegistrationResponse{
		Id:           user.Id,
		AccessToken:  "access",
		RefreshToken: "refresh",
	}, nil
}

func (h *Handler) Login(ctx context.Context, in *tgstorev1.LoginRequest) (*tgstorev1.LoginResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
