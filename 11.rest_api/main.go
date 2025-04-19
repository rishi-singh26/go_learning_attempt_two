package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rishisingh.in/event-manager/db"
	"rishisingh.in/event-manager/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Could not featch events"})
		return
	}

	ctx.JSON(http.StatusOK, events)
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
