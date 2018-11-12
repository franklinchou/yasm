package models

type UserRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type LogoutRequest struct {
	SessionId string `form:"session-id" json:"session-id"`
}
