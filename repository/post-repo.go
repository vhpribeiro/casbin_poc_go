package repository

import (
	"go_tutorial_post.com/entity"
)

type IPostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]*entity.Post, error)
}
