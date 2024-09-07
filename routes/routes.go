package routes

import (
	"booking-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// creating a route group so we can apply the Authenticate middleware to all routes in this group
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", register)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// a way to add a middleware/filter to a single route
	//server.POST("/events", middleware.Authenticate, createEvent)

	server.POST("/signup", registerUser)
	server.POST("/login", login)
}
