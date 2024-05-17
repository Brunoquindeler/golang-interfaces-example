package infra

import "github.com/brunoquindeler/golang-interfaces-example/internal/entity"

type UserStoreI interface {
	Create(user *entity.User)
}
