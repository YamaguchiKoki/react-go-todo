package domain

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type (
	User struct {
		ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Name     string             `json:"name" bson:"name"`
		Email    string             `json:"email" bson:"email"`
		Password string             `json:"-" bson:"password"`
	}

	UserRepository interface {
		CreateUser(ctx context.Context, user *User) error
		GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)
		GetUserByEmail(ctx context.Context, email string) (*User, error)
		UpdateUser(ctx context.Context, user *User) error
		DeleteUser(ctx context.Context, id primitive.ObjectID) error
	}
)

func NewUser(id primitive.ObjectID, name, email, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) SetPassword(password string) {
	u.Password = password
}


func (u *User) IsPasswordValid(password string) bool {
	return u.Password == password
}
