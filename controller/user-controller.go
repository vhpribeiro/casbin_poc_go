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
	AddRoleForUserInDomain(response http.ResponseWriter, request *http.Request)
	GetTheRolesFromAUserInDomain(response http.ResponseWriter, request *http.Request)
	AddPolicy(response http.ResponseWriter, request *http.Request)
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
	user := request.Header.Get("user")
	domain := GetTheSingleKey(response, request, "domain")
	resource := GetTheSingleKey(response, request, "resource")
	action := GetTheSingleKey(response, request, "action")

	if user == "" || domain == "" || resource == "" || action == "" {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting some value"})
		return
	}

	if userService.CheckIfUserHasPermission(user, domain, resource, action) {
		response.WriteHeader(http.StatusOK)
		return
	}

	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: "Usuario não possui permisão"})
}

func (*userController) AddRoleForUserInDomain(response http.ResponseWriter, request *http.Request) {
	var userRoleDomainDto dtos.UserRoleDomainDto
	err := json.NewDecoder(request.Body).Decode(&userRoleDomainDto)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the request"})
		return
	}

	if userService.AddRoleForUserInDomain(userRoleDomainDto.User, userRoleDomainDto.Role, userRoleDomainDto.Domain) {
		response.WriteHeader(http.StatusOK)
		return
	}

	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: "Erro não foi possível adicionar o papel ao usuário"})
}

func (*userController) GetTheRolesFromAUserInDomain(response http.ResponseWriter, request *http.Request) {
	var rolesFromUser []string
	user := request.Header.Get("user")
	domain := GetTheSingleKey(response, request, "domain")

	if user == "" || domain == "" {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting some value"})
		return
	}

	rolesFromUser = userService.GetTheRolesFromAUserInDomain(user, domain)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(rolesFromUser)
	return
}

func (*userController) AddPolicy(response http.ResponseWriter, request *http.Request) {
	var policyDto dtos.PolicyDto
	err := json.NewDecoder(request.Body).Decode(&policyDto)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the request"})
		return
	}

	if userService.AddPolicy(policyDto.Role, policyDto.Domain, policyDto.Resource, policyDto.Action) {
		response.WriteHeader(http.StatusOK)
		return
	}

	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: "Erro não foi possível adicionar o papel ao usuário"})
}

func GetTheSingleKey(response http.ResponseWriter, request *http.Request, keyToGet string) string {
	keys, ok := request.URL.Query()[keyToGet]

	if !ok || len(keys[0]) < 1 {
		return ""
	}

	return keys[0]
}
