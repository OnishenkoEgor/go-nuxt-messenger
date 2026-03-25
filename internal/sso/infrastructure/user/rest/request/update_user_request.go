package request

type UpdateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
