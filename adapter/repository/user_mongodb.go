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

//domain層で定義したrepositoryのインターフェースの実装
//TODO:シグネチャ合わせるとこから
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
		ID: user.ID,
		
	}
}

func (u UserNoSQL) FindByID(ctx context.Context, id primitive.ObjectID) (domain.User, error) {
	
}

func (u UserNoSQL) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	userBSON := &userBSON{
		ID: user.ID,

	}
}

func (u UserNoSQL) Update(ctx context.Context, user domain.User) error {
	userBSON := &userBSON{
		ID: user.ID,

	}
}

func (u UserNoSQL) Delete(ctx context.Context, user domain.User) error {
	userBSON := &userBSON{
		ID: user.ID,

	}
}