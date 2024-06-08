package repositories

import (
    "context"
    "todo-app/internal/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
    collection *mongo.Collection
}

func NewTodoRepository(db *models.DB) *TodoRepository {
    return &TodoRepository{
        collection: db.Client.Database("todoapp").Collection("todos"),
    }
}

func (r *TodoRepository) Add(todo *models.Todo) error {
    _, err := r.collection.InsertOne(context.Background(), todo)
    return err
}

func (r *TodoRepository) GetById(id string) (*models.Todo, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var todo models.Todo
    err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&todo)
    return &todo, err
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
    cursor, err := r.collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var todos []models.Todo
    for cursor.Next(context.Background()) {
        var todo models.Todo
        cursor.Decode(&todo)
        todos = append(todos, todo)
    }

    return todos, cursor.Err()
}

func (r *TodoRepository) UpdateById(id string, update *models.Todo) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = r.collection.UpdateOne(
        context.Background(),
        bson.M{"_id": objID},
        bson.M{"$set": update},
    )
    return err
}

func (r *TodoRepository) DeleteById(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = r.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
    return err
}
