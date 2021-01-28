package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"go_tutorial_post.com/repository"
)

type IPolicyService interface {
	AddPolicy(role string, domain string, resource string, action string, attribute string) bool
}

type policyService struct{}

var (
	cRepo           repository.ICasbinRepository
	cMongoDbAdapter persist.BatchAdapter
	enf             *casbin.Enforcer
	serviceError    error
)

func NewPolicyService(casbinRepository repository.ICasbinRepository) IPolicyService {
	cRepo = casbinRepository
	cMongoDbAdapter = casbinRepository.GetTheAdapter()
	enf, serviceError = casbin.NewEnforcer("./casbin/rbac_with_domains_model.conf", casbinMongoDbAdapter)
	if serviceError != nil {
		panic(serviceError)
	}
	return &policyService{}
}

func (*policyService) AddPolicy(role string, domain string, resource string, action string, attribute string) bool {
	result, errs := enforce.AddPolicy(role, domain, resource, action, attribute)
	if errs != nil {
		panic(errs)
	}

	return result
}
