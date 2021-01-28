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
	httpRouter     router.IRouter             = router.NewChiRouter()
	postRepository repository.IPostRepository = repository.NewMongoPostRepository()
	postService    service.IPostService       = service.NewPostService(postRepository)
	postController controller.IPostController = controller.NewPostController(postService)
	userService    service.IUserService       = service.NewUserService()
	userController controller.IUserController = controller.NewUserController(userService)
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and Runnning")
	})

	httpRouter.GET("/posts", postController.GetPosts)

	httpRouter.POST("/posts", postController.AddPosts)

	//OS DADOS QUE TO PEGANDO NO BODY AQUI, DEVERIA VIR TUDO PELA URL, DEPOIS TROCA
	httpRouter.GET("/users", userController.CheckIfUserHasPermission)

	//DEPOIS TROCA ISSO POR UM PUT
	httpRouter.POST("/users", userController.AddRoleToUser)

	httpRouter.SERVE(port)
}
