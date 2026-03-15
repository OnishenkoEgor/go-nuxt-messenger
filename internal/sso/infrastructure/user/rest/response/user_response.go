package response

import (
	"messenger/sso/domain/user"
)

type UserResponse struct {
	Id       int    `json:"id"`
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
