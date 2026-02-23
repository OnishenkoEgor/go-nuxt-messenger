package response

import (
	"sso/domain/user"
)

type UserResponse struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewUserResponse(user *user.User) *UserResponse {
	return &UserResponse{
		Id:       user.GetId(),
		Login:    user.GetLogin(),
		Password: user.GetPassword(),
	}
}
