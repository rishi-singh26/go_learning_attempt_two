package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rishisingh.in/event-manager/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Not authorozed."})
		return
	}

	ctx.Set("userId", userId)

	ctx.Next()
}
