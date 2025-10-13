package auth

import (
	"context"

	ssov1 "github.com/k3rlll/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPCServer *grpc.Server) {
	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{})

}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest) (
	*ssov1.LoginResponse, error) {
	panic("Implement me")
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest) (
	*ssov1.RegisterResponse, error) {
	panic("Implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest) (
	*ssov1.IsAdminResponse, error) {
	panic("Implement me")
}
