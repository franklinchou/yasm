package main

import (
	"github.com/gin-gonic/gin"
	"./handlers"
	"fmt"
)

//*********************************************************
// Application defaults
//*********************************************************
const ApplicationDefaultPort = 9000

func main() {
	router := gin.Default()
	router.GET("/health", handlers.HealthCheckHandler)
	router.GET("/health/redis", handlers.HealthRedisHandler)
	router.POST("/session", handlers.CreateSessionHandler)

	// Start the server
	router.Run(fmt.Sprintf(":%d", ApplicationDefaultPort))
}