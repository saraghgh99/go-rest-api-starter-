package main

import (
	"rest-api-pra/db"
	"rest-api-pra/routes"
    "github.com/gin-gonic/gin"
)

// @title EventGo API
// @version 1.0
// @description This is a REST API for an Event Management System.
// @host go-rest-api-starter.onrender.com
// @BasePath /

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
