package routes

import (
	"api_example/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id."})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventById(id)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not complete request. Try again later."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not complete request. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func unregisterFromEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id."})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventById(id)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not complete request. Try again later."})
		return
	}

	err = event.Unregister(userId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not complete request. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered!"})
}
