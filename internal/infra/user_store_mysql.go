package infra

import (
	"fmt"

	"github.com/brunoquindeler/golang-interfaces-example/internal/entity"
)

type MySQLUserStore struct{}

func NewMySQLUserStore() *MySQLUserStore {
	return &MySQLUserStore{}
}

func (s *MySQLUserStore) Create(user *entity.User) {
	fmt.Printf("User %q created in MySQL", user.Name)
}
