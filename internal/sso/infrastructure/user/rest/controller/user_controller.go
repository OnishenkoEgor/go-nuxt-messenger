package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sso/application"
	command2 "sso/application/user/command"
	query2 "sso/application/user/query"
	"sso/infrastructure/user/rest/request"
	"sso/infrastructure/user/rest/response"
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
	users, err := c.app.Queries.GetUsersQuery.Handle(query2.GetUsersQuery{})

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
	user, err := c.app.Queries.GetUserByIdQuery.Handle(query2.GetUserByIdQuery{
		Id: "qwe",
	})

	if err != nil {
		WriteResponse[string](true, "Error on get user", w)
		fmt.Println(err)

		return
	}

	userResponse := response.NewUserResponse(user)

	WriteResponse[*response.UserResponse](false, userResponse, w)
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateUserRequest
	err := ReadRequest[request.CreateUserRequest](r, &userRequest)

	if err != nil {
		panic(err)
	}

	err = c.app.Commands.CreateUserCommand.Handle(command2.CreateUserCommand{
		Login:    userRequest.Login,
		Password: userRequest.Password,
	})

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

	err := c.app.Commands.DeleteUserCommand.Handle(command2.DeleteUserCommand{Id: id})

	if err != nil {
		WriteResponse[string](true, "Error on delete user", w)
		fmt.Println(err)

		return
	}

	WriteResponse[interface{}](false, nil, w)
}
