// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"goTgTemplate/app/connections"
	"goTgTemplate/internal/controller"
	"goTgTemplate/internal/middleware"
	"goTgTemplate/internal/repository"
	"goTgTemplate/internal/service"
)

// Injectors from wire.go:

func Init() *Initialization {
	db := connections.ConnectToDB()
	testRepositoryImpl := repository.TestRepositoryInit(db)
	testServiceImpl := service.TestServiceInit(testRepositoryImpl)
	testControllerImpl := controller.TestControllerInit(testServiceImpl)
	middlewaresImpl := middleware.MiddlewaresInit()
	initialization := NewInitialization(testControllerImpl, middlewaresImpl)
	return initialization
}

// wire.go:

var connectionsSet = wire.NewSet(connections.GetBotInstance, connections.ConnectToDB, connections.ConnectToNatsBroker, connections.ConnectToRedis)

var testSet = wire.NewSet(repository.TestRepositoryInit, wire.Bind(new(repository.TestRepository), new(*repository.TestRepositoryImpl)), service.TestServiceInit, wire.Bind(new(service.TestService), new(*service.TestServiceImpl)), controller.TestControllerInit, wire.Bind(new(controller.TestController), new(*controller.TestControllerImpl)))

var workSet = wire.NewSet(middleware.MiddlewaresInit, wire.Bind(new(middleware.Middlewares), new(*middleware.MiddlewaresImpl)))
