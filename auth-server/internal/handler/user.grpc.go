package handler

import (
	"context"
	"grpc-auth-app/auth-server/api"
	"grpc-auth-app/auth-server/internal/domain"
	"grpc-auth-app/auth-server/pkg/config"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type AuthServer struct {
	api.UnimplementedUserServiceServer
	userSvc domain.UserService
	router  *gin.Engine
}

func NewAuthServer(svc domain.UserService, r *gin.Engine) *AuthServer {
	return &AuthServer{
		userSvc: svc,
		router:  r,
	}
}

func (a *AuthServer) GetUser(ctx context.Context, req *api.UserRequest) (*api.User, error) {
	return a.userSvc.SigninService(ctx, req)
}
func (a *AuthServer) CreateUser(ctx context.Context, req *api.User) (*api.CreateUserResponse, error) {
	return a.userSvc.SignUpService(ctx, req)
}
func (a *AuthServer) ListUsers(ctx context.Context, req *api.ListUserRequest) (*api.ListUserResponse, error) {
	return a.userSvc.ListUsersService(ctx, req)
}
func (a *AuthServer) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	return a.userSvc.DeleteUserService(ctx, req)
}
func RegisterServer(grpcServer *grpc.Server, svcServer *AuthServer) {
	api.RegisterUserServiceServer(grpcServer, svcServer)
}
func StartGrpcServer(svc *AuthServer) {
	grpcPort := config.Config.GRPCPort
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	RegisterServer(grpcServer, svc)
	log.Printf("Server is running on port %s\n", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
