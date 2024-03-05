package appgrpc

import (
	"context"
	"log/slog"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/olegtemek/tg-store/internal/service"
	"github.com/olegtemek/tg-store/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ServiceAuthFuncOverride interface {
	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
}
type Interceptors struct {
	log         *slog.Logger
	accessToken string
	userService service.User
}

func NewInterceptors(log *slog.Logger, accessToken string, userService *service.User) *Interceptors {
	return &Interceptors{
		log:         log,
		accessToken: accessToken,
		userService: *userService,
	}
}

func (i *Interceptors) InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		i.log.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (i *Interceptors) InterceptorAuth(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	newCtx, err := i.authFunc(ctx, info.FullMethod, i.accessToken)
	if err != nil {
		return nil, err
	}
	return handler(newCtx, req)
}

func (i *Interceptors) authFunc(ctx context.Context, method string, accessTokenSecret string) (context.Context, error) {
	if method == "/tgstore.UserService/Login" || method == "/tgstore/UserService/Registration" {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	values := md["authorization"]

	if len(values) < 1 {
		return ctx, status.Errorf(codes.Unauthenticated, "cannot get access token")
	}
	authorization := strings.Split(values[0], " ")

	if len(authorization) < 1 {
		return ctx, status.Errorf(codes.Unauthenticated, "cannot get access token")
	}

	userId, err := utils.CompareToken(authorization[1], accessTokenSecret)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid")
	}

	user, wrappErr := i.userService.GetById(userId)
	if wrappErr != nil {
		return ctx, status.Errorf(codes.Unauthenticated, wrappErr.Error())
	}

	ctx = context.WithValue(ctx, "user", user)

	return ctx, nil
}
