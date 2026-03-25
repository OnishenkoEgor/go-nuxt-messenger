package repository

import (
	"database/sql"
	"errors"
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
	_, err := repo.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2);", user.GetLogin(), user.GetPassword())

	return err
}

func (repo UserRepository) GetById(id int) (*user.User, error) {
	rows, err := repo.db.Query("SELECT * FROM users WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}

	users, err := repo.parseRows(rows)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("not found user by id")
	}

	return users[0], nil
}

func (repo UserRepository) Get() ([]*user.User, error) {
	rows, err := repo.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}

	return repo.parseRows(rows)
}

func (repo UserRepository) Update(user *user.User) error {
	_, err := repo.db.Exec("UPDATE users SET login = $1, password = $2 WHERE id= $3", user.GetLogin(), user.GetPassword(), user.GetId())

	return err
}

func (repo UserRepository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM users WHERE id = $1;", id)

	return err
}

func (repo UserRepository) parseRows(rows *sql.Rows) ([]*user.User, error) {
	var users []*user.User

	for rows.Next() {
		var id int
		var login string
		var password string

		err := rows.Scan(&id, &login, &password)

		var u = &user.User{}

		u.SetId(id)
		u.SetLogin(login)
		u.SetPassword(password)

		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
