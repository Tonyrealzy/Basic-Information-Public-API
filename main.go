package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Tonyrealzy/Basic-Information-Public-API/handlers"
	"github.com/Tonyrealzy/Basic-Information-Public-API/middleware"
)

func main () {
	r := gin.Default()

	// Apply CORS middleware
	r.Use(middleware.SetUpCORS())

	// Define routes
	r.GET("/info", handlers.GetBasicInfo)

	// Start the server
	r.Run(":8080")
}