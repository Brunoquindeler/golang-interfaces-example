package service_test

import (
	"testing"

	"github.com/brunoquindeler/golang-interfaces-example/internal/entity"
	"github.com/brunoquindeler/golang-interfaces-example/internal/service"
)

type MockUserStore struct {
	users []*entity.User
}

func (m *MockUserStore) Create(user *entity.User) {
	m.users = append(m.users, user)
}

func TestCreateUserHandler(t *testing.T) {
	userName := "Teste"

	mockUserStore := &MockUserStore{
		users: []*entity.User{},
	}

	userService := service.NewUserService(mockUserStore)

	userService.CreateUserHandler(userName)

	if mockUserStore.users[0] == nil {
		t.Fatal("user is nil")
	}

	if mockUserStore.users[0].Name != userName {
		t.Fatal("user not found")
	}
}
