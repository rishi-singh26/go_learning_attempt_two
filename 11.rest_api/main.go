package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rishisingh.in/event-manager/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
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

	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{"status": true, "message": "Event created", "event": event})
}
