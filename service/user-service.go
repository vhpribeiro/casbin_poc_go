package service

import (
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
)

type IUserService interface {
	AddRoleForUserInDomain(user string, domain string, role string) bool
	CheckIfUserHasPermission(user string, domain string, resource string, action string) bool
	GetTheRolesFromAUserInDomain(user string, domain string) []string
	AddPolicy(role string, domain string, resource string, action string) bool
}

type userService struct{}

var (
	connectionString string = "localhost:27017"
)

func NewUserService() IUserService {
	return &userService{}
}

func (*userService) AddRoleForUserInDomain(user string, role string, domain string) bool {
	adapter, err1 := mongodbadapter.NewAdapter("localhost:27017") // Your MongoDB URL.
	if err1 != nil {
		panic(err1)
	}

	enforce, err2 := casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", adapter)
	if err2 != nil {
		panic(err2)
	}

	result, err2 := enforce.AddRoleForUserInDomain(user, role, domain)
	if err2 != nil {
		panic(err2)
	}
	return result
}

func (*userService) CheckIfUserHasPermission(user string, domain string, resource string, action string) bool {
	adapter, err1 := mongodbadapter.NewAdapter("localhost:27017") // Your MongoDB URL.
	if err1 != nil {
		panic(err1)
	}

	enforce, err2 := casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", adapter)
	if err2 != nil {
		panic(err2)
	}

	result, err2 := enforce.Enforce(user, domain, resource, action)
	if err2 != nil {
		panic(err2)
	}
	return result
}

func (*userService) GetTheRolesFromAUserInDomain(user string, domain string) []string {
	adapter, err1 := mongodbadapter.NewAdapter("localhost:27017") // Your MongoDB URL.
	if err1 != nil {
		panic(err1)
	}

	enforce, err2 := casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", adapter)
	if err2 != nil {
		panic(err2)
	}

	return enforce.GetRolesForUserInDomain(user, domain)
}

func (*userService) AddPolicy(role string, domain string, resource string, action string) bool {

	adapter, err1 := mongodbadapter.NewAdapter("localhost:27017") // Your MongoDB URL.
	if err1 != nil {
		panic(err1)
	}

	enforce, err2 := casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", adapter)
	if err2 != nil {
		panic(err2)
	}

	result, err3 := enforce.AddPolicy(role, domain, resource, action)
	if err3 != nil {
		panic(err3)
	}

	return result
}
