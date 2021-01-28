package service

import (
	"github.com/casbin/casbin"
)

type IUserService interface {
	AddRoleForUser(user string, domain string, role string) bool
	CheckIfUserHasPermission(user string, domain string, resource string, action string) bool
}

type userService struct{}

var (
	enforce = casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", "./casbin/rbac_with_domains_policy.csv")
)

func NewUserService() IUserService {
	return &userService{}
}

func (*userService) AddRoleForUser(user string, role string, domain string) bool {
	if enforce.AddRoleForUserInDomain(user, role, domain) {
		return true
	}
	return false
}

func (*userService) CheckIfUserHasPermission(user string, domain string, resource string, action string) bool {
	return enforce.Enforce(user, domain, resource, action)
}
