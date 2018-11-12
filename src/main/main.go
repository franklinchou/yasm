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
	router.GET("/sessions", handlers.SessionGetHandler)
	router.GET("/session/:sessionId", handlers.SessionValidateHandler)
	router.POST("/session", handlers.SessionCreateHandler)
	router.DELETE("/session/:sessionId", handlers.SessionInvalidateHandler)

	// Start the server
	router.Run(fmt.Sprintf(":%d", ApplicationDefaultPort))
}
