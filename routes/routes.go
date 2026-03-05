package routes

import (
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {
	server.GET("/events", getEvents)    // handler for an incoming GET request
	server.GET("/events/:id", getEvent) // /events/1, /events/5

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", loginUser)
}
