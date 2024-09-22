package domain

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrTodoNotFound = errors.New("Todo not found")
)

type (
	Todo struct {
		ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
		Body string
		Completed bool
	}

	TodoRepository interface {
		CreateTodo(ctx context.Context, todo *Todo) error
		GetTodosByUserID(ctx context.Context, userID primitive.ObjectID) ([]*Todo, error)
		UpdateTodo(ctx context.Context, todo *Todo) error
		DeleteTodo(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error
	}
)

func NewTodo(id primitive.ObjectID, userID primitive.ObjectID, body string) *Todo {
	return &Todo {
		ID: id,
		UserID: userID,
		Body: body,
		Completed: false,
	}
}
