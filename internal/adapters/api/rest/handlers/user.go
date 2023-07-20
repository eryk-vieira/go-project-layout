package handlers

import (
	"net/http"

	"github.com/eryk-vieira/go-api-project-layout/internal/core/dto"
	"github.com/eryk-vieira/go-api-project-layout/internal/core/port"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

func (service *UserHandler) CreateUser(c *gin.Context) {
	var newUser dto.CreateUser

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	user, err := service.UserService.Create(&newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, user)

	return
}
