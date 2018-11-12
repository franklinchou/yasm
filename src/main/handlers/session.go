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

func SessionCreateHandler(ctx *gin.Context) {
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

func SessionGetHandler(ctx *gin.Context) {
	if !DebugMode {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "Unauthorized"})
		return
	}
	keyValues, _ := services.GetSessions(DefaultLimit)
	ctx.JSON(http.StatusOK, gin.H{"data": keyValues})
}

func SessionValidateHandler(ctx *gin.Context) {
	sessionId := ctx.Param("sessionId")
	token, err := services.GetTokenBySession(sessionId)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func SessionInvalidateHandler(ctx *gin.Context) {
	sessionId := ctx.Param("sessionId")
	e := services.DeleteSession(sessionId)
	if e != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":    false,
			"session-id": sessionId,
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
