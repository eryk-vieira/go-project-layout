package rest

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/api/rest/handlers"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres/repository"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/port"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/service"
	"github.com/gin-gonic/gin"
)

var (
	userService port.UserService
)

func init() {
	userRepository := repository.NewUserRepository(postgres.GetDatabaseInstance())
	userService = service.NewUserService(userRepository)
}

func RunServer() {
	router := gin.Default()

	v1 := router.Group("/v1")

	appHandler := handlers.NewAppHandler()
	v1.GET("/health-check", appHandler.HealthCheck)

	userHandler := handlers.NewUserHandler(userService)
	users := v1.Group("/users")
	users.POST("/users", userHandler.CreateUser)

	router.Run(":8080")
}
