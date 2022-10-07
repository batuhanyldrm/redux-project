package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/greetings/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func (repository *Repository) CreateTodo(todo models.Todo) error {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, todo)

	if err != nil {
		return err
	}

	return nil

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

func (repository *Repository) GetTodos() ([]models.Todo, error) {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	todos := []models.Todo{}
	for cur.Next(ctx) {
		var todo models.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		//go'da ekleme append ile yapılır
		todos = append(todos, todo)
	}

	return todos, nil

}

func (repository *Repository) PostTodos(todo models.Todo) error {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, todo)

	if err != nil {
		return err
	}

	return nil

}

func (repository *Repository) DeleteTodos(todoId string) error {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": todoId}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}

func (repository *Repository) UpdateTodos(todo models.Todo, ID string) error {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := collection.FindOneAndReplace(ctx, bson.M{"id": ID}, todo)

	if result == nil {
		return result.Err()
	}

	/* updatedTodo, err := repository.GetTodo(ID)

	if err != nil {
		return nil, err
	} */

	return nil

}

func (repository *Repository) GetTodo(ID string) (models.Todo, error) {
	collection := repository.client.Database("todo").Collection("todo")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todo := models.Todo{}
	err := collection.FindOne(ctx, bson.M{}).Decode(&todo); 
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("dsafdfsafd", todo)

	return todo, nil

}

func GetCleanTestRepository() *Repository {

	repository := NewRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	todoDB := repository.client.Database("todo")
	todoDB.Drop(ctx)

	return repository
}
