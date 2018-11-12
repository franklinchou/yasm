package models


type UserRequest struct {
	Email	string `form:"email" json:"email" binding:"required"`
}