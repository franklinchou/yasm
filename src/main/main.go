package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"./handlers"
	"./services"
	"github.com/gomodule/redigo/redis"
)

//*********************************************************
// Application defaults
//*********************************************************
const ApplicationDefaultPort = 9000

// Provide a redis connection pool
var MyPool *redis.Pool = services.NewPool()

func main() {

	// Inject connection pool into handlers
	handlers.MyPool = MyPool
	services.MyPool = MyPool

	// Inject debug mode
	handlers.DebugMode = true // TODO Get from env var

	router := gin.Default()
	router.GET("/health", handlers.HealthCheckHandler)
	router.GET("/health/redis", handlers.HealthRedisHandler)
	router.GET("/sessions", handlers.GetSessionsHandler)
	router.GET("/session/:sessionId", handlers.ValidateSessionHandler)
	router.POST("/session", handlers.CreateSessionHandler)
	router.DELETE("/session/:sessionId", handlers.InvalidateSessionHandler)

	// Start the server
	router.Run(fmt.Sprintf(":%d", ApplicationDefaultPort))
}
