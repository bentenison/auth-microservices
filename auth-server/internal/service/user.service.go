package service

import (
	"context"
	"grpc-auth-app/auth-server/api"
	"grpc-auth-app/auth-server/internal/domain"
	"grpc-auth-app/auth-server/pkg/auth"
	"log"
)

type USConf struct {
	UserRepo domain.UserRepository
	TokenSvc *auth.Auth
}
type userService struct {
	userRepo domain.UserRepository
	tokenSvc *auth.Auth
}

func NewUserService(conf *USConf) domain.UserService {
	return &userService{
		userRepo: conf.UserRepo,
		tokenSvc: conf.TokenSvc,
	}
}
func (us *userService) SigninService(ctx context.Context, req *api.UserRequest) (*api.User, error) {
	id := req.GetUserId()
	user, err := us.userRepo.GetUser(id)
	if err != nil {
		log.Println("error occured while signing in", err)
		return nil, err
	}
	// log.Println("got id", id)
	return user, nil
}
func (us *userService) SignUpService(ctx context.Context, req *api.User) (*api.CreateUserResponse, error) {
	// log.Println("user is ", req.Guid)
	user, err := us.userRepo.CreateUser(req)
	if err != nil {
		log.Println("error occured while signing up user", err)
		return nil, err
	}
	token, err := us.tokenSvc.CreateToken(user.Id)
	if err != nil {
		log.Println("error occured while creating user token", err)
		return nil, err
	}
	// log.Println("returning token", token)
	return &api.CreateUserResponse{Token: &api.Token{Token: token}}, nil
}
func (us *userService) DeleteUserService(ctx context.Context, req *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	id := req.GetUserId()
	res, err := us.userRepo.DeleteUser(id)
	if err != nil {
		log.Println("error occured while deleting user", err)
		return nil, err
	}
	// out:= &api.User{
	// 	Id: res.ID,
	// }
	return &api.DeleteUserResponse{User: &api.User{Id: res.Id}}, nil
}
func (us *userService) ListUsersService(ctx context.Context, req *api.ListUserRequest) (*api.ListUserResponse, error) {
	users, err := us.userRepo.ListUsers()
	if err != nil {
		log.Println("error occured while listing users", err)
		return nil, err
	}
	// res := []*api.User.(users)
	return &api.ListUserResponse{Users: users}, nil
}
