//go:build wireinject
// +build wireinject

package main

import (
	"gochicoba/handler"
	"gochicoba/producer"
	"gochicoba/repository"
	"gochicoba/service"

	"github.com/go-redis/redis"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func ItemHandler(db *gorm.DB, redis *redis.Client) handler.ItemHandler {
	wire.Build(repository.NewItemRepository, service.NewItemService, handler.NewItemHandler)
	return handler.ItemHandler{}
}

func UserHandler(db *gorm.DB) handler.UserHandler {
	wire.Build(repository.NewUserRepository, service.NewUserService, handler.NewUserHandler)
	return handler.UserHandler{}
}

func BuyHandler(db *gorm.DB, redis *redis.Client, mb *producer.MessageBroker) handler.BuyHandler {
	wire.Build(repository.NewBuyRepository, repository.NewItemRepository, repository.NewUserRepository, service.NewUserService, service.NewItemService, service.NewBuyService, handler.NewBuyHandler)
	return handler.BuyHandler{}
}

func LoginHandler(db *gorm.DB) handler.LoginHandler {
	wire.Build(repository.NewLoginRepository, service.NewLoginService, handler.NewLoginHandler)
	return handler.LoginHandler{}
}

// func LoginHandler(db *gorm.DB) handler.LoginHandler {
// 	wire.Build(repository.NewLoginRepository, service.NewLoginService, handler.NewLoginHandler)
// 	return handler.LoginHandler{}
// }
