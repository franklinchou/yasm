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
	token := ctx.Param("token")
	sessionId, err := services.GetSessionByToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"session-id": sessionId})
}

func SessionInvalidateHandler(ctx *gin.Context) {
	token := ctx.Param("token")
	e := services.DeleteSession(token)
	if e != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":    false,
			"user-token": token,
			"error":      e.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
