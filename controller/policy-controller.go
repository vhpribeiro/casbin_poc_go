package controller

import (
	"encoding/json"
	"net/http"

	"go_tutorial_post.com/errors"
	"go_tutorial_post.com/service"
	"go_tutorial_post.com/service/dtos"
)

type IPolicyController interface {
	AddPolicy(response http.ResponseWriter, request *http.Request)
}

type policyController struct{}

var (
	policyService service.IPolicyService
)

func NewPolicyController(service service.IPolicyService) IPolicyController {
	policyService = service
	return &policyController{}
}

func (*policyController) AddPolicy(response http.ResponseWriter, request *http.Request) {
	var policyDto dtos.PolicyDto
	err := json.NewDecoder(request.Body).Decode(&policyDto)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the request"})
		return
	}

	if policyService.AddPolicy(policyDto.Role, policyDto.Domain, policyDto.Resource, policyDto.Action) {
		response.WriteHeader(http.StatusOK)
		return
	}

	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: "Erro não foi possível adicionar o papel ao usuário"})
}
