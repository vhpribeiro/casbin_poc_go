package service

import (
	"github.com/casbin/casbin"
)

type IUserService interface {
	AddRoleForUserInDomain(user string, domain string, role string) bool
	CheckIfUserHasPermission(user string, domain string, resource string, action string) bool
	GetTheRolesFromAUserInDomain(user string, domain string) []string
	AddPolicy(role string, domain string, resource string, action string) bool
}

type userService struct{}

var (
	enforce = casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", "./casbin/rbac_with_domains_policy.csv")
)

func NewUserService() IUserService {
	return &userService{}
}

func (*userService) AddRoleForUserInDomain(user string, role string, domain string) bool {
	if enforce.AddRoleForUserInDomain(user, role, domain) {
		return true
	}
	return false
}

func (*userService) CheckIfUserHasPermission(user string, domain string, resource string, action string) bool {
	return enforce.Enforce(user, domain, resource, action)
}

func (*userService) GetTheRolesFromAUserInDomain(user string, domain string) []string {
	return enforce.GetRolesForUserInDomain(user, domain)
}

func (*userService) AddPolicy(role string, domain string, resource string, action string) bool {
	return enforce.AddPolicy(role, domain, resource, action)
}
