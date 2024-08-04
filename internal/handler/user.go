package handler

import (
	"net/http"

	service "github.com/DavidEsdrs/go-mercado/internal/services"
	"github.com/DavidEsdrs/go-mercado/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	log     *logger.Logger
	service *service.UserService
}

func NewUserHandler(service *service.UserService, log *logger.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		log:     log,
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
	c.Status(http.StatusOK)
}

func (uh *UserHandler) Signup(c *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to parse body",
		})
		c.Abort()
		return
	}

	user, err := uh.service.Signup(body.Username, body.Email, body.Password)
	if err != nil {
		uh.log.Info("status 400 - bad request: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"created_at": user.CreatedAt.String(),
		"updated_at": user.UpdatedAt.String(),
	})
}
