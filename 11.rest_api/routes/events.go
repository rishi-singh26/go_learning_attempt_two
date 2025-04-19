package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rishisingh.in/event-manager/models"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not featch events"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not featch events"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not featch event"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not parse request data."})
		return
	}
	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not create event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": true, "message": "Event created", "event": event})
}
