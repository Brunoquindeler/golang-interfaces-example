package main

import (
	"github.com/brunoquindeler/golang-interfaces-example/internal/infra"
	"github.com/brunoquindeler/golang-interfaces-example/internal/service"
)

func main() {
	// userStore := infra.NewMySQLUserStore()
	userStore := infra.NewSQLiteUserStore()

	userService := service.NewUserService(userStore)

	userService.CreateUserHandler("Bruno")
}
