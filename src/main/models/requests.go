package models

type CreateSessionRequest struct {
	Token string `form:"token" json:"token"`
}

type UserRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type LogoutRequest struct {
	SessionId string `form:"session-id" json:"session-id"`
}
