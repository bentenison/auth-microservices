package domain

import (
	"context"
	"grpc-auth-app/auth-server/api"
)

type Friend struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID            string   `json:"_id"`
	Index         int      `json:"index"`
	GUID          string   `json:"guid"`
	IsActive      bool     `json:"isActive"`
	Balance       string   `json:"balance"`
	Picture       string   `json:"picture"`
	Age           int      `json:"age"`
	EyeColor      string   `json:"eyeColor"`
	Name          string   `json:"name"`
	Gender        string   `json:"gender"`
	Company       string   `json:"company"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	About         string   `json:"about"`
	Registered    string   `json:"registered"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Tags          []string `json:"tags"`
	Friends       []Friend `json:"friends"`
	Greeting      string   `json:"greeting"`
	FavoriteFruit string   `json:"favoriteFruit"`
}

type UserRepository interface {
	GetUser(id string) (*api.User, error)
	CreateUser(*api.User) (*api.User, error)
	DeleteUser(id string) (*api.User, error)
	ListUsers() ([]*api.User, error)
}
type UserService interface {
	SigninService(context.Context, *api.UserRequest) (*api.User, error)
	SignUpService(context.Context, *api.User) (*api.CreateUserResponse, error)
	DeleteUserService(context.Context, *api.DeleteUserRequest) (*api.DeleteUserResponse, error)
	ListUsersService(context.Context, *api.ListUserRequest) (*api.ListUserResponse, error)
}
