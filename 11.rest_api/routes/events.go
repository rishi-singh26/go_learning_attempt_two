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

	userId := ctx.GetInt64("userId")
	event.UserId = userId

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not create event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": true, "message": "Event created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not update event"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not update event"})
		return
	}

	if event.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Not authorized to update event"})
		return
	}

	var updateEvent models.Event

	err = ctx.ShouldBindJSON(&updateEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not parse request data."})
		return
	}

	updateEvent.ID = eventId

	err = updateEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not update event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Event updated successfully", "event": updateEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Could not delete event"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Event not found"})
		return
	}

	if event.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Not authorized to delete event"})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not delete event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Event deleted successfully"})
}
