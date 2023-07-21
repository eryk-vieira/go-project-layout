package main

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/api/rest"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/api/rest/handlers"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres/repository"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/service"
)

func main() {
	appHandler := handlers.NewAppHandler()
	userHandler := injectUserHandlerDependencies()

	rest.RunServer(appHandler, userHandler)
}

func injectUserHandlerDependencies() *handlers.UserHandler {
	userRepository := repository.NewUserRepository(postgres.GetDatabaseInstance())
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	return userHandler
}
