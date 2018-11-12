package models

type CreateSessionRequest struct {
	Token string `form:"token" json:"token"`
}
