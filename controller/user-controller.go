package controller

import (
	"encoding/json"
	"net/http"

	"go_tutorial_post.com/errors"
	"go_tutorial_post.com/service"
	"go_tutorial_post.com/service/dtos"
)

type IUserController interface {
	CheckIfUserHasPermission(response http.ResponseWriter, request *http.Request)
	AddRoleToUser(response http.ResponseWriter, request *http.Request)
}

type userController struct{}

var (
	userService service.IUserService
)

func NewUserController(service service.IUserService) IUserController {
	userService = service
	return &userController{}
}

func (*userController) CheckIfUserHasPermission(response http.ResponseWriter, request *http.Request) {
	var userObjectAction dtos.UserObjectActionDto
	err := json.NewDecoder(request.Body).Decode(&userObjectAction)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the request"})
		return
	}

	if userService.CheckIfUserHasPermission(userObjectAction.User, userObjectAction.Object, userObjectAction.Action) {
		response.WriteHeader(http.StatusOK)
		return
	}

	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: "Usuario não possui permisão"})
}

func (*userController) AddRoleToUser(response http.ResponseWriter, request *http.Request) {
	var userRole dtos.UserRoleDto
	err := json.NewDecoder(request.Body).Decode(&userRole)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the request"})
		return
	}

	if userService.AddRoleForUser(userRole.User, userRole.Role) {
		response.WriteHeader(http.StatusOK)
		return
	}

	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: "Erro não foi possível adicionar o papel ao usuário"})
}
