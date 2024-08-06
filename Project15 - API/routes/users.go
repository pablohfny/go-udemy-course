package routes

import (
	"api_example/models"
	"api_example/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request Formation."})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request Formation."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Signup Successful", "event": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request Formation."})
		return
	}

	err = user.Authenticate()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials."})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error generating token."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Signup Successful", "token": token})
}
