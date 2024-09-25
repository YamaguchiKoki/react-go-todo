package repository

import (
	"context"

	"github.com/YamaguchiKoki/react-go-todo/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userBSON struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

//domain層で定義したrepositoryのインターフェースの実装
//db操作はinterface/mongo_handler.go
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

func (u UserNoSQL) Create(ctx context.Context, user *domain.User) error {
	//TODO:encrypt
	userBSON := &userBSON{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	if err := u.db.Store(ctx, u.collectionName, userBSON); err != nil {
		return errors.Wrap(err, "error creating user")
	}

	return nil
}

func (u UserNoSQL) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	var user *domain.User

	err := u.db.FindOne(ctx, u.collectionName, bson.M{"_id": id}, nil, &user)
	if err != nil {
		return &domain.User{}, errors.Wrap(err, "error finding user")
	}

	return user, nil
}

func (u UserNoSQL) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (u UserNoSQL) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (u UserNoSQL) Delete(ctx context.Context, id primitive.ObjectID) error {
	return nil
}