package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rishisingh.in/event-manager/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Event not found"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not register user for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": true, "message": "Registered"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not parse event id"})
		return
	}

	var event models.Event

	event.ID = eventId
	err = event.CancelRegistration(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not cancel registration"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Registration cancelled"})
}
