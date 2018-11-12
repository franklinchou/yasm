package handlers

import (
	"net/http"
	"../services"
	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func HealthRedisHandler(ctx *gin.Context) {
	_, e := services.Connect().Ping().Result()
	if e != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": e.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	}
}
