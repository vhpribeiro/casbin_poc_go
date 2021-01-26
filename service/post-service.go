package service

import (
	"errors"
	"math/rand"

	"go_tutorial_post.com/entity"
	"go_tutorial_post.com/repository"
)

type IPostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]*entity.Post, error)
}

var (
	repo repository.IPostRepository
)

type service struct{}

func NewPostService(repository repository.IPostRepository) IPostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]*entity.Post, error) {
	return repo.FindAll()
}
