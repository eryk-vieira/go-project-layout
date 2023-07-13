package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (*AppHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
	})

	return
}
