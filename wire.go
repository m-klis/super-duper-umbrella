//go:build wireinject
// +build wireinject

package main

import (
	"gochicoba/handler"
	"gochicoba/repository"
	"gochicoba/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func ItemHandler(db *gorm.DB) handler.ItemHandler {
	wire.Build(repository.NewItemRepository, service.NewItemService, handler.NewItemHandler)
	return handler.ItemHandler{}
}

func UserHandler(db *gorm.DB) handler.UserHandler {
	wire.Build(repository.NewUserRepository, service.NewUserService, handler.NewUserHandler)
	return handler.UserHandler{}
}

func BuyHandler(db *gorm.DB) handler.BuyHandler {
	wire.Build(repository.NewBuyRepository, service.NewBuyService, handler.NewBuyHandler)
	return handler.BuyHandler{}
}

// func LoginHandler(db *gorm.DB) handler.LoginHandler {
// 	wire.Build(repository.NewLoginRepository, service.NewLoginService, handler.NewLoginHandler)
// 	return handler.LoginHandler{}
// }
