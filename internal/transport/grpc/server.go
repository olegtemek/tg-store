package appgrpc

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/olegtemek/tg-store/internal/config"
	"github.com/olegtemek/tg-store/internal/service"
	"github.com/olegtemek/tg-store/internal/transport/grpc/user"
	tgstorev1 "github.com/olegtemek/tg-store/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	Log         *slog.Logger
	Cfg         *config.Config
	Srv         *grpc.Server
	UserHandler tgstorev1.UnimplementedUserServiceServer
}

func New(log *slog.Logger, cfg *config.Config, services *service.Service) *Server {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	interceptors := NewInterceptors(log, cfg.AccessTokenSecret, &services.User)

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(interceptors.InterceptorLogger(), loggingOpts...),
		grpc.UnaryServerInterceptor(interceptors.InterceptorAuth),
	))

	// INIT ALL SERVICES
	tgstorev1.RegisterUserServiceServer(server, user.NewGRPCHandler(log, &services.User, cfg.AccessTokenSecret, cfg.RefreshTokenSecret))

	return &Server{
		Log: log,
		Cfg: cfg,
		Srv: server,
	}
}

func (s *Server) ListenAndServe() error {
	addr := fmt.Sprintf(":%s", s.Cfg.GRPC.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("server: %w", err)
	}

	s.Log.Info("grpc server started", slog.String("addr", lis.Addr().String()))

	if err := s.Srv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() {
	s.Log.Info("stopping gRPC server")
	s.Srv.GracefulStop()
}
