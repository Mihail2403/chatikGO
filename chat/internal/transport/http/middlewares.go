package handler

import (
	"chat/internal/server"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenSlice := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		if len(tokenSlice) != 2 {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			log.Println("Give token on wrong format: \"Bearer <TOKEN>\"")
			return
		}

		token := tokenSlice[1]
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			log.Println("Token is empty")
			return
		}

		request, err := http.NewRequest("GET", server.AUTH_URL+"/auth/get-by-token", nil)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {

			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			log.Println("Token is invalid")
			return
		}
		var authResponse struct {
			Data struct {
				Id int `json:"id"`
			} `json:"data"`
		}
		err = json.NewDecoder(response.Body).Decode(&authResponse)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			log.Println(err)
			return
		}
		ctx.Set("user_id", authResponse.Data.Id)
		ctx.Next()
	}
}
