package rest

import (
	"github.com/eryk-vieira/go-api-project-layout/internal/adapters/api/rest/handlers"
	"github.com/gin-gonic/gin"
)

func RunServer(
	appHandler *handlers.AppHandler,
	userHandler *handlers.UserHandler,
) {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/health-check", appHandler.HealthCheck)

	users := v1.Group("/users")
	users.POST("/users", userHandler.CreateUser)

	router.Run(":8080")
}
