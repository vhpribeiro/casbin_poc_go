package main

import (
	"fmt"
	"net/http"

	"go_tutorial_post.com/controller"
	router "go_tutorial_post.com/http"
	"go_tutorial_post.com/repository"
	"go_tutorial_post.com/service"
)

var (
	httpRouter       router.IRouter               = router.NewChiRouter()
	casbinRepository repository.ICasbinRepository = repository.NewCasbinMongoRepository()

	userService   service.IUserService   = service.NewUserService(casbinRepository)
	policyService service.IPolicyService = service.NewPolicyService(casbinRepository)

	userController   controller.IUserController   = controller.NewUserController(userService)
	policyController controller.IPolicyController = controller.NewPolicyController(policyService)
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and Runnning")
	})

	httpRouter.GET("/users", userController.CheckIfUserHasPermission)

	//DEPOIS TROCA ISSO POR UM PUT
	httpRouter.POST("/users/roles", userController.AddRoleForUserInDomain)

	httpRouter.GET("/users/roles", userController.GetTheRolesFromAUserInDomain)

	httpRouter.POST("/policy", policyController.AddPolicy)

	httpRouter.SERVE(port)
}
