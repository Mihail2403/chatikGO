package handler

import (
	"auth/internal/service"

	"github.com/gin-gonic/gin"
)

// handler struct
type Handler struct {
	service *service.Service
}

// Constructor Handler
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Register gin router
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.GET("/get-by-token", h.GetUserByToken)
	}

	return router
}
