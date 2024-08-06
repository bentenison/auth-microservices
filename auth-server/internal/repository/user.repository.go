package repository

import (
	"fmt"
	"grpc-auth-app/auth-server/api"
	"grpc-auth-app/auth-server/internal/domain"
	"log"
)

type UserReposoitory struct {
}

type userRepo struct {
	users map[string]*api.User
}

func NewUserRepository() domain.UserRepository {
	return &userRepo{
		users: make(map[string]*api.User),
	}
}

func (u *userRepo) GetUser(id string) (*api.User, error) {
	user, ok := u.users[id]
	if !ok {
		return nil, fmt.Errorf("error no user found with id %s", id)
	}
	return user, nil
}

func (u *userRepo) CreateUser(usr *api.User) (*api.User, error) {
	u.users[usr.Id] = usr
	return usr, nil
}
func (u *userRepo) DeleteUser(id string) (*api.User, error) {
	user, err := u.GetUser(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	delete(u.users, id)
	return user, nil
}
func (u *userRepo) ListUsers() ([]*api.User, error) {
	users := make([]*api.User, len(u.users))
	// lists := []*api.User{}
	for _, v := range u.users {
		users = append(users, v)
		// lists = append(lists, v)
	}
	return users, nil
}
