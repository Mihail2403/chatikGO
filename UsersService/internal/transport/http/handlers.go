package handler

import (
	"users-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("", h.GetAllUsers)
		users.POST("", h.CreateUser)
		users.GET("/by-many-id", h.GetUsersByManyId)
		users.GET("/:id", h.GetById)
		users.PATCH("/:id", h.UpdateUserById)
		users.DELETE("/:id", h.DeleteUser)
	}
	return r
}
