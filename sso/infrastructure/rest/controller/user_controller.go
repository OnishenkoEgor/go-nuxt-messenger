package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sso/application"
	"sso/application/query"
	"sso/infrastructure/rest/request"
	"sso/infrastructure/rest/response"
)

type Controller struct {
	app application.Application
}

func NewUserController(app application.Application) Controller {
	return Controller{
		app: app,
	}
}

func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.app.Queries.GetUsersQuery.Handle(query.GetUsersQuery{})

	if err != nil {
		WriteResponse[string](true, "Error on get users", w)
	}

	var responseUsersList []*response.UserResponse

	for _, user := range users {
		responseUsersList = append(responseUsersList, response.NewUserResponse(user))
	}

	WriteResponse[[]*response.UserResponse](false, responseUsersList, w)
}

func (c Controller) Get(w http.ResponseWriter, r *http.Request) {
	user, err := c.service.GetById("qwr")

	if err != nil {
		WriteResponse[string](true, "Error on get user", w)
		fmt.Println(err)

		return
	}

	userResponse := response.NewUserResponse(user)

	WriteResponse[*response.UserResponse](false, userResponse, w)
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	var user request.CreateUserRequest
	err := ReadRequest[request.CreateUserRequest](r, &user)

	if err != nil {
		panic(err)
	}

	err = c.service.Create(user.Login, user.Password)

	if err != nil {
		WriteResponse[string](true, "Error on create user", w)
		fmt.Println(err)
		return
	}

	WriteResponse[any](false, nil, w)
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request) {

}

func (c Controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		WriteResponse[string](true, "Error on delete user", w)
		return
	}

	err := c.service.Delete(id)

	if err != nil {
		WriteResponse[string](true, "Error on delete user", w)
		fmt.Println(err)

		return
	}

	WriteResponse[interface{}](false, nil, w)
}
