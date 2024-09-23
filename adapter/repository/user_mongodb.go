package repository

import (
	"context"

	"github.com/YamaguchiKoki/react-go-todo/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userBSON struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserNoSQL struct {
	collectionName string
	db NoSQL
}

func NewUserNoSQL(db NoSQL) UserNoSQL {
	return UserNoSQL{
		db: db,
		collectionName: "users",
	}
}

func (u UserNoSQL) Create(ctx context.Context, user domain.User) (domain.User, error) {
	userBSON := &userBSON{
		ID: ,
	}
}