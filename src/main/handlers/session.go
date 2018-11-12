package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../services"
)

//*********************************************************
// Defaults
//*********************************************************
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

func ValidateSessionHandler(ctx *gin.Context) {
	sessionId := ctx.Param("session-id")
	token, err := services.GetTokenBySession(sessionId)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func InvalidateSessionHandler(ctx *gin.Context) {
	sessionId := ctx.Param("session-id")
	e := services.DeleteSession(sessionId)
	if e != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":    false,
			"session-id": sessionId,
		})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{})
	}

}
