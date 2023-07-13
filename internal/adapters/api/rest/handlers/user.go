package handlers

import (
	"net/http"

	"github.com/eryk-vieira/go-api-project-layout/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		service: userService,
	}
}

func (*UserHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
	})

	return
}
