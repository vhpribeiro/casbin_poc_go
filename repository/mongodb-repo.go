package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_tutorial_post.com/entity"
)

type repoMongo struct{}

var (
	collection         *mongo.Collection
	ctx                       = context.TODO()
	collectionPostName string = "posts"
	databaseName       string = "post_database"
	connectionString   string = "mongodb://localhost:27017"
)

//Construtor
func NewMongoPostRepository() IPostRepository {
	return &repoMongo{}
}

func (*repoMongo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to create a Mongo Client: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(databaseName).Collection(collectionPostName)

	_, err = collection.InsertOne(ctx, &post)

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	//Retornar nil, indica que NÃO teve erros
	return post, nil
}

func (*repoMongo) FindAll() ([]*entity.Post, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to create a Frestore Client: %v", err)
		return nil, err
	}

	filter := bson.D{{}}
	var posts []*entity.Post

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return posts, err
	}

	for cur.Next(ctx) {
		var p entity.Post
		err := cur.Decode(&p)
		if err != nil {
			return posts, err
		}

		posts = append(posts, &p)
	}

	if err := cur.Err(); err != nil {
		return posts, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)
	defer client.Disconnect(ctx)

	if len(posts) == 0 {
		return posts, mongo.ErrNoDocuments
	}

	return posts, nil
}
