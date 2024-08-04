package handler

import (
	"net/http"

	service "github.com/DavidEsdrs/go-mercado/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		c.Abort()
		return
	}

	tokenString, err := uh.service.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized request",
		})
		c.Abort()
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*2, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}
