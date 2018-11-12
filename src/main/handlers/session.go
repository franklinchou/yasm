package handlers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"net/http"
)

func CreateSessionHandler(ctx *gin.Context) {
	var req models.CreateSessionRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Unauthorized"})
		return
	}

	sessionId := createSession("", req.token)
	ctx.JSON(http.StatusOK, gin.H{"session_id": sessionId})
}
