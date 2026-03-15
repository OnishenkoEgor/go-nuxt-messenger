package repository

import (
	"database/sql"
	"fmt"
	"messenger/sso/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) Create(user *user.User) error {
	fmt.Println("test2")
	_, err := repo.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2);", user.GetLogin(), user.GetPassword())
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) GetById(id string) (*user.User, error) {
	var u user.User

	rows, err := repo.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&u)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo UserRepository) Get() ([]*user.User, error) {
	var users []*user.User

	rows, err := repo.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var login string
		var password string

		err = rows.Scan(&id, &login, &password)

		var u = user.User{}

		u.SetId(id)
		u.SetLogin(login)
		u.SetPassword(password)

		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

func (repo UserRepository) Update(user *user.User) error {

	return nil
}

func (repo UserRepository) Delete(id string) error {

	return nil
}
