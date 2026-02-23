package repository

import (
	"database/sql"
	"sso/domain/user"
)

const (
	dbName         string = "sso"
	collectionName string = "users"
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
	//_, err := repo.collection.InsertOne(*repo.ctx, user)

	//if err != nil {
	//	return err
	//}

	return nil
}

func (repo UserRepository) GetById(id string) (*user.User, error) {
	//var user user.User
	//
	//filter := bson.D{{"uuid", id}}
	//err := repo.collection.FindOne(*repo.ctx, filter).Decode(&user)
	//
	//if err != nil {
	//return nil, err
	//}
	//
	//return &user, nil
	return nil, nil
}

func (repo UserRepository) Get() ([]*user.User, error) {
	//cursor, err := repo.collection.Find(*repo.ctx, bson.D{})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//var users []*user.User
	//err = cursor.All(*repo.ctx, &users)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return users, nil

	return nil, nil
}

func (repo UserRepository) Update(user *user.User) error {
	return nil
}

func (repo UserRepository) Delete(id string) error {
	//filter := bson.D{{"uuid", bson.D{{"$eq", id}}}}
	//fmt.Println(id)
	//
	//cursor, err := repo.collection.Find(*repo.ctx, filter)
	//if err != nil {
	//	return err
	//}
	//
	//var res []user.User
	//
	//err = cursor.All(context.TODO(), &res)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(res)
	//
	//result, err := repo.collection.DeleteOne(*repo.ctx, filter)
	//fmt.Println(result)
	//if err != nil {
	//	return err
	//}
	//
	//if result.DeletedCount != 1 {
	//	fmt.Println(result.DeletedCount)
	//	return errors.New("incorrect removed rows count")
	//}

	return nil
}
