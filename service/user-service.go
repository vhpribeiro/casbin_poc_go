package service

import (
	"github.com/casbin/casbin"
)

type IUserService interface {
	AddRoleForUser(user string, role string) bool
	CheckIfUserHasPermission(user string, object string, action string) bool
}

type userService struct{}

var (
	enforce = casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", "./casbin/rbac_with_domains_policy.csv")
)

func NewUserService() IUserService {
	return &userService{}
}

func (*userService) AddRoleForUser(user string, role string) bool {
	if enforce.AddRoleForUser(user, role) {
		return true
	}
	return false
}

func (*userService) CheckIfUserHasPermission(user string, object string, action string) bool {
	return enforce.Enforce(user, object, action)
}
