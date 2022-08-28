package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/greetings/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	client *mongo.Client
}

func (repository *Repository) GetTodo(ID string) (models.Todo, error) {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todo := models.Todo{}

	if err := collection.FindOne(ctx, bson.M{"id": ID}).Decode(&todo); err != nil {
		log.Fatal(err)
	}

	return todo, nil
}

func (repository *Repository) PostTodos(todos models.Todo) (models.Todo, error) {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// burada  insertone ın 3 tane isteği olduğu için options girdim sanırım olmasa da olur farkı var mı
	_, err := collection.InsertOne(ctx, todos, options.InsertOne())

	if err != nil {
		return models.Todo{}, err
	}
	createdTodo, _ := repository.GetTodo(todos.ID)
	fmt.Println(createdTodo, "sdsad")
	return createdTodo, nil
}

func GetCleanTestRepository() *Repository {

	repository := NewTestRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	todoDB := repository.client.Database("todos")
	todoDB.Drop(ctx)

	return repository
}

func NewRepository() *Repository {
	uri := "mongodb+srv://Cluster:bthn998877@cluster0.hnmuy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func NewTestRepository() *Repository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}
