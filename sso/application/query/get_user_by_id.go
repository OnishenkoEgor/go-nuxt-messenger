package query

import "sso/domain/user"

type GetUserByIdQuery struct {
	id string
}

type GetUserByIdQueryHandler struct {
	repo user.Repository
}

func NewGetUserByIdQueryHandler(repo user.Repository) GetUserByIdQueryHandler {
	return GetUserByIdQueryHandler{
		repo: repo,
	}
}

func (h GetUserByIdQueryHandler) Handle(q GetUserByIdQuery) (*user.User, error) {
	return h.repo.GetById(q.id)
}
