package main

import (
	"net/http"
	"rest-api/database"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()
 // /events is the route to getEvents
	server.GET("/events", getEvents) // handler for an incoming GET request
	server.GET("/events/:id", getEvent) // /events/1, /events/5
	server.POST("/events", createEvents)

	server.Run(":8080") // starts the server with localhost  
}

func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}
	context.JSON(http.StatusOK, events) // means that everything worked 
	// The H is shortcut for creating a map
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get a path paramether value
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // works like a print func

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events.", "cause" : err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}