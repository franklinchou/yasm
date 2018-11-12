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

	// Inject the pool into handlers
	handlers.MyPool = MyPool
	services.MyPool = MyPool

	router := gin.Default()
	router.GET("/health", handlers.HealthCheckHandler)
	router.GET("/health/redis", handlers.HealthRedisHandler)
	router.POST("/session", handlers.CreateSessionHandler)

	// Start the server
	router.Run(fmt.Sprintf(":%d", ApplicationDefaultPort))
}
