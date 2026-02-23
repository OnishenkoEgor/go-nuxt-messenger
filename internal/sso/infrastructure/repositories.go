package infrastructure

import (
	"database/sql"
	"sso/domain"
	"sso/infrastructure/user/repository"
)

func NewRepositories(dbClient *sql.DB) domain.Repositories {
	userRepo := repository.NewUserRepository(dbClient)

	return domain.Repositories{
		UserRepository: userRepo,
	}
}
