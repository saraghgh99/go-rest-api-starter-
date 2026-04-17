package routes

import (
	"rest-api-pra/middlewares"
	"github.com/gin-gonic/gin"

	// 1. Add these three Swagger imports:
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "rest-api-pra/docs" // This imports the docs folder you just generated!
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /events/5

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

	// 2. Add this ONE line to serve the generated UI:
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}