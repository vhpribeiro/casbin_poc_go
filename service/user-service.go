package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"go_tutorial_post.com/repository"
)

type IUserService interface {
	AddRoleForUserInDomain(user string, domain string, role string, attribute string) bool
	CheckIfUserHasPermission(user string, domain string, resource string, action string, attribute string) bool
	GetTheRolesFromAUserInDomain(user string, domain string, attribute string) []string
}

type userService struct{}

var (
	casbinRepo           repository.ICasbinRepository
	casbinMongoDbAdapter persist.BatchAdapter
	enforce              *casbin.Enforcer
	err                  error
)

func NewUserService(casbinRepository repository.ICasbinRepository) IUserService {
	casbinRepo = casbinRepository
	casbinMongoDbAdapter = casbinRepository.GetTheAdapter()
	enforce, err = casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", casbinMongoDbAdapter)
	if err != nil {
		panic(err)
	}
	return &userService{}
}

func (*userService) AddRoleForUserInDomain(user string, domain string, role string, attribute string) bool {
	result, errs := enforce.AddGroupingPolicy(user, role, domain, attribute)
	if errs != nil {
		panic(errs)
	}
	return result
}

func (*userService) CheckIfUserHasPermission(user string, domain string, resource string, action string, attribute string) bool {
	result, errs := enforce.Enforce(user, domain, resource, action, attribute)
	if errs != nil {
		panic(errs)
	}
	return result
}

func (*userService) GetTheRolesFromAUserInDomain(user string, domain string, attribute string) []string {
	result, errs := enforce.GetRolesForUser(user, domain, attribute)
	if errs != nil {
		panic(errs)
	}
	return result
}
