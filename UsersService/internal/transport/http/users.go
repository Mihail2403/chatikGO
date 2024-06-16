package handler

import (
	"log"
	"strconv"
	"users-service/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	var input entity.User
	err := ctx.BindJSON(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "invalid body"})
		return
	}
	err = h.service.Users.Create(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "error creating user"})
		return
	}
	ctx.JSON(201, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetAllUsers(ctx *gin.Context) {
	users, err := h.service.Users.GetAll()
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "error getting users"})
		return
	}
	ctx.JSON(200, gin.H{"users": users})
}

func (h *Handler) GetUsersByManyId(ctx *gin.Context) {
	var input struct {
		Ids []int `json:"ids"`
	}
	err := ctx.BindJSON(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "invalid body"})
		return
	}
	users, err := h.service.Users.GetByIDArray(input.Ids)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "error getting users"})
		return
	}
	ctx.JSON(200, gin.H{"users": users})
}

func (h *Handler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.service.Users.GetByID(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, gin.H{"user": user})
}

func (h *Handler) UpdateUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	var input entity.User
	err = ctx.BindJSON(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "invalid body"})
		return
	}
	err = h.service.Users.Update(id, &input)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "error updating user"})
		return
	}
	ctx.JSON(200, gin.H{"message": "user updated successfully"})
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	err = h.service.Users.Delete(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "error deleting user"})
		return
	}
	ctx.JSON(200, gin.H{"message": "user deleted successfully"})
}
