package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rishisingh.in/event-manager/models"
	"rishisingh.in/event-manager/utils"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": true, "message": "User created"})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerageToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Login Successful", "token": token})
}
