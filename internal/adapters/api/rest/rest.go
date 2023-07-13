package rest

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/api/rest/handlers"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres"
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/database/postgres/repository"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	v1 := router.Group("/v1")

	appHandler := handlers.NewAppHandler()
	v1.GET("/health-check", appHandler.HealthCheck)

	userHandler := handlers.NewUserHandler(service.NewUserService(repository.NewUserRepository(postgres.GetDatabaseInstance())))
	users := v1.Group("/users")
	users.POST("/users", userHandler.CreateUser)
}
