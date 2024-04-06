package dto

import (
	"github.com/ak-karimzai/web-labs/internal/model"
	"time"
)

// UserInfo
// @Description User info
type UserInfo struct {
	FirstName string    `json:"first_name" example:"User first name"`
	LastName  string    `json:"last_name" example:"User last name"`
	Username  string    `json:"username" example:"User's username'"`
	CreateAt  time.Time `json:"create_at" example:"2023-04-01T15:04:05Z"`
}

// LoginResponse
// @Description Login request response
type LoginResponse struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"user_info"`
}

func NewLoginResponse(token string, userInfo model.User) *LoginResponse {
	return &LoginResponse{Token: token, UserInfo: UserInfo{
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Username:  userInfo.Username,
		CreateAt:  userInfo.CreatedAt,
	}}
}
