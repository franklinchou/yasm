package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../services"
)

const DefaultLimit = 100

//*********************************************************
// Models
//*********************************************************
type createSessionRequest struct {
	Token string `form:"token" json:"token"`
}

//*********************************************************

func CreateSessionHandler(ctx *gin.Context) {
	var req createSessionRequest
	err := ctx.BindJSON(&req)
	noToken := req.Token == ""

	if err != nil || noToken {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Unauthorized"})
		return
	}

	token, sessionId := services.CreateSession(req.Token)
	ctx.JSON(http.StatusOK, gin.H{
		"token":      token,
		"session-id": sessionId,
	})
}

func GetSessionsHandler(ctx *gin.Context) {
	if !DebugMode {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "Unauthorized"})
		return
	}
	keyValues, _ := services.GetSessions(DefaultLimit)
	ctx.JSON(http.StatusOK, gin.H{"data": keyValues})
}
