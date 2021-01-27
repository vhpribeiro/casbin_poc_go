package controller

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
)

type IUserController interface {
	CheckIfUserHasPermission(response http.ResponseWriter, request *http.Request)
}

type userController struct{}

func NewUserController() IUserController {
	return &userController{}
}

func (*userController) CheckIfUserHasPermission(response http.ResponseWriter, request *http.Request) {
	//Ta pegando certo
	enforce := casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", "./casbin/rbac_with_domains_policy.csv")

	// dom := "domain1" // the domain
	sub := "alice"  // the user that wants to access a resource.
	obj := "data1"  // the resource that is going to be accessed.
	act := "custom" // the operation that the user performs on the resource.

	if res := enforce.Enforce(sub, obj, act); res {
		fmt.Printf("Deu bom")

	} else {
		fmt.Printf("Deu ruim")
	}
}
