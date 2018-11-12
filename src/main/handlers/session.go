package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../models"
	"../services"
)


func CreateSessionHandler(ctx *gin.Context) {
	var req models.CreateSessionRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Unauthorized"})
		return
	}

	sessionId := services.CreateSession(req.Token)
	ctx.JSON(http.StatusOK, gin.H{"session-id": sessionId})
}
