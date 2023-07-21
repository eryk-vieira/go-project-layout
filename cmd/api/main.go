package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/api/rest"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/api/rest/handlers"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres/repository"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/service"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		err := godotenv.Load()

		fmt.Println("Trying to load .env file")

		if err != nil {
			log.Panic("Error trying to load .env file")
		}
	}

	postgres.CreateDatabaseInstance()

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
