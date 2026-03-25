package query

import (
	"messenger/sso/domain/user"
)

type GetUsersQuery struct {
}

type GetUsersQueryHandler struct {
	repo user.Repository
}

func NewGetUsersQueryHandler(repo user.Repository) GetUsersQueryHandler {
	return GetUsersQueryHandler{
		repo: repo,
	}
}

func (h *GetUsersQueryHandler) Handle(_ GetUsersQuery) ([]*user.User, error) {
	users, err := h.repo.Get()

	return users, err
}
