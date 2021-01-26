package controller

import (
	"encoding/json"
	"net/http"

	"go_tutorial_post.com/entity"
	"go_tutorial_post.com/errors"
	"go_tutorial_post.com/service"
)

type IPostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPosts(response http.ResponseWriter, request *http.Request)
}

var (
	postService service.IPostService
)

type controller struct{}

func NewPostController(service service.IPostService) IPostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the post"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPosts(response http.ResponseWriter, request *http.Request) {
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the request"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
