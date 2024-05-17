package service

import (
	"github.com/brunoquindeler/golang-interfaces-example/internal/entity"
	"github.com/brunoquindeler/golang-interfaces-example/internal/infra"
)

type UserService struct {
	userStore infra.UserStoreI
}

func NewUserService(userStore infra.UserStoreI) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

func (s *UserService) CreateUserHandler(name string) {
	user := entity.NewUser(name)
	s.userStore.Create(user)
}
