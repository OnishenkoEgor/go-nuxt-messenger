package query

import "messenger/sso/domain/user"

type GetUserByIdQuery struct {
	Id int
}

type GetUserByIdQueryHandler struct {
	repo user.Repository
}

func NewGetUserByIdQueryHandler(repo user.Repository) GetUserByIdQueryHandler {
	return GetUserByIdQueryHandler{
		repo: repo,
	}
}

func (h *GetUserByIdQueryHandler) Handle(q GetUserByIdQuery) (*user.User, error) {
	return h.repo.GetById(q.Id)
}
