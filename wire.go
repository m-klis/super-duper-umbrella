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
	wire.build(repository.NewUserRepository, service.NewUserService, handler.NewUserHandler)
	return handler.ItemHandler{}
}
