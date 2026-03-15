package infrastructure

import (
	"database/sql"
	"messenger/sso/domain"
	"messenger/sso/infrastructure/user/repository"
)

func NewRepositories(dbClient *sql.DB) domain.Repositories {
	userRepo := repository.NewUserRepository(dbClient)

	return domain.Repositories{
		UserRepository: userRepo,
	}
}
