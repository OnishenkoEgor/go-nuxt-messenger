package controller

import (
	"fmt"
	"messenger/router"
	"messenger/sso/application"
	"messenger/sso/application/user/command"
	"messenger/sso/application/user/query"
	"messenger/sso/infrastructure/user/rest/request"
	"messenger/sso/infrastructure/user/rest/response"
	"net/http"
	"strconv"
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
		router.WriteResponse[string](true, "Error on get users", w)
		fmt.Println(err)
		return
	}

	var responseUsersList []*response.UserResponse

	for _, user := range users {
		responseUsersList = append(responseUsersList, response.NewUserResponse(user))
	}

	router.WriteResponse[[]*response.UserResponse](false, responseUsersList, w)
}

func (c Controller) Get(w http.ResponseWriter, r *http.Request) {
	vars := router.Variables(r)
	id, ok := vars["id"]

	if !ok {
		router.WriteResponse[string](true, "Error on get user", w)
		return
	}

	preparedId, err := strconv.Atoi(id)
	if err != nil {
		router.WriteResponse[string](true, "Error: id is not INTEGER", w)
	}

	user, err := c.app.Queries.GetUserByIdQuery.Handle(query.GetUserByIdQuery{
		Id: preparedId,
	})

	if err != nil {
		router.WriteResponse[string](true, "Error on get user", w)
		fmt.Println(err)

		return
	}

	userResponse := response.NewUserResponse(user)

	router.WriteResponse[*response.UserResponse](false, userResponse, w)
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateUserRequest
	err := router.ParseRequest[request.CreateUserRequest](r, &userRequest)

	if err != nil {
		router.WriteResponse[string](true, "Error: failed parse request data", w)
		return
	}

	if userRequest.Login == "" || userRequest.Password == "" {
		router.WriteResponse[string](true, "Error: user data empty", w)
		return
	}

	err = c.app.Commands.CreateUserCommand.Handle(command.CreateUserCommand{
		Login:    userRequest.Login,
		Password: userRequest.Password,
	})

	if err != nil {
		router.WriteResponse[string](true, "Error on create user", w)
		fmt.Println(err)
		return
	}

	router.WriteResponse[any](false, nil, w)
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request) {
	vars := router.Variables(r)
	id, ok := vars["id"]
	if !ok {
		router.WriteResponse[string](true, "Error on delete user", w)
		return
	}

	preparedId, err := strconv.Atoi(id)
	if err != nil {
		router.WriteResponse[string](true, "Id is not INTEGER", w)
		return
	}

	var userRequest request.UpdateUserRequest
	err = router.ParseRequest[request.UpdateUserRequest](r, &userRequest)

	if err != nil {
		router.WriteResponse[string](true, "Failed parse request data", w)
		return
	}

	if userRequest.Login == "" || userRequest.Password == "" {
		router.WriteResponse[string](true, "User data empty", w)
		return
	}

	err = c.app.Commands.UpdateUserCommand.Handle(command.UpdateUserCommand{
		Id:       preparedId,
		Login:    userRequest.Login,
		Password: userRequest.Password,
	})

	if err != nil {
		router.WriteResponse[string](false, "Failed update user", w)
		return
	}

	router.WriteResponse[interface{}](false, nil, w)
}

func (c Controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := router.Variables(r)
	id, ok := vars["id"]
	if !ok {
		router.WriteResponse[string](true, "Error on delete user", w)
		return
	}

	preparedId, err := strconv.Atoi(id)
	if err != nil {
		router.WriteResponse[string](true, "Error: id is not INTEGER", w)
	}

	err = c.app.Commands.DeleteUserCommand.Handle(command.DeleteUserCommand{Id: preparedId})

	if err != nil {
		router.WriteResponse[string](true, "Error on delete user", w)
		fmt.Println(err)

		return
	}

	router.WriteResponse[interface{}](false, nil, w)
}
