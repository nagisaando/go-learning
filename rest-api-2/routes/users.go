package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api-2/models"
	"example.com/rest-api-2/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = user.ValidateCredential()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
		return
	}

	fmt.Println("id ", user.ID)
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login success!", "token": token})
}
