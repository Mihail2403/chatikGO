package handler

import (
	"auth/entity"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// handler  for user register
func (h *Handler) SignUp(c *gin.Context) {

	// Binding request to struct
	usrReq := entity.UserFromJSON{}
	if err := c.ShouldBindJSON(&usrReq); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		return
	}

	// binding from request struct to entity
	user, err := usrReq.ToEntity()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err})
		return
	}

	// start signing up
	if err := h.service.SignUp(user); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		return
	}

	// response
	c.JSON(201, gin.H{})
}

// handler for login
func (h *Handler) SignIn(c *gin.Context) {

	// Binding  request to struct
	usrReq := entity.UserFromJSON{}
	if err := c.ShouldBindJSON(&usrReq); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		return
	}

	// binding from request struct to entity
	user, err := usrReq.ToEntity()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(400, gin.H{"error": err})
		return
	}

	// getting the token from  service
	token, err := h.service.GetToken(user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// response
	c.JSON(200, gin.H{"token": token})
}

// get  user info by token
func (h *Handler) GetUserByToken(c *gin.Context) {
	// get  token from header
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// getting User from Service
	user, err := h.service.GetUserByToken(tokenString)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}
	// response
	c.JSON(200, gin.H{"data": user})
}
