package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"go_tutorial_post.com/entity"
)

type repo struct{}

//Construtor
func NewFirestorePostRepository() IPostRepository {
	return &repo{}
}

const (
	projectId      string = "golangrestapi"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Frestore Client: %v", err)
		return nil, err
	}

	//Essa linha é executada depois que o método inteiro é feito
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	//Retornar nil, indica que NÃO teve erros
	return post, nil
}

func (*repo) FindAll() ([]*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	//Essa linha é executada depois que o método inteiro é feito
	defer client.Close()

	var posts []*entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, &post)
	}

	return posts, nil
}
