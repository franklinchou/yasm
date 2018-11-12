package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"./handlers"
	"./services"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

//*********************************************************
// Application defaults
//*********************************************************
const ApplicationDefaultPort = 9000

var DebugMode bool

func init() {
	debugStr := os.Getenv("DEBUG_MODE")

	if debugStr == "" {
		err := fmt.Errorf("yasm: debug mode must be set")
		panic(err)
	}

	DebugMode, _ = strconv.ParseBool(debugStr)
	return
}

//*********************************************************

// Provide a redis connection pool
var MyPool *redis.Pool = services.NewPool()

func main() {

	// Inject connection pool into handlers
	handlers.MyPool = MyPool
	services.MyPool = MyPool

	// Inject debug mode
	handlers.DebugMode = DebugMode

	router := gin.Default()
	router.GET("/health", handlers.HealthCheckHandler)
	router.GET("/health/redis", handlers.HealthRedisHandler)
	router.GET("/sessions", handlers.SessionGetHandler)
	router.GET("/session/:token", handlers.SessionValidateHandler)
	router.POST("/session", handlers.SessionCreateHandler)
	router.DELETE("/session/:token", handlers.SessionInvalidateHandler)

	// Start the server
	router.Run(fmt.Sprintf(":%d", ApplicationDefaultPort))
}
