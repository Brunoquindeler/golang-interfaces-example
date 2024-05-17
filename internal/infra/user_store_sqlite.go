package infra

import (
	"fmt"

	"github.com/brunoquindeler/golang-interfaces-example/internal/entity"
)

type SQLiteUserStore struct{}

func NewSQLiteUserStore() *SQLiteUserStore {
	return &SQLiteUserStore{}
}

func (s *SQLiteUserStore) Create(user *entity.User) {
	fmt.Printf("User %q created in SQLite", user.Name)
}
