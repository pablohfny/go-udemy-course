package middlewares

import (
	"api_example/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing Authorization Header."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
