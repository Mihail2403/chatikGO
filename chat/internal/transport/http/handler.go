package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// handler struct
type Handler struct {
}

// Constructor Handler
func NewHandler() *Handler {
	return &Handler{}
}

// Register gin routerc
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	auth.Use(AuthMiddleware())
	{
		auth.GET("/sign-up", func(ctx *gin.Context) {
			fmt.Println("JJJJJ")
			user_id := 0
			if id, ok := ctx.Get("user_id"); ok {
				user_id = id.(int)
			}
			ctx.JSON(200, gin.H{"id": user_id})
		})
	}

	return router
}
