package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/YamaguchiKoki/react-go-todo/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrUserAlreadyExists = errors.New("user with this email already exists")
)
type (
	CreateUserUsecase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Name string `json:"name" validete:"required"`
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	CreateUserPresenter interface {
		Output(domain.User) CreateUserOutput
	}

	CreateUserOutput struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
	}

	CreateUserInteractor struct {
		repo domain.UserRepository
		presenter CreateUserPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateUserInteractor(
	repo domain.UserRepository,
	presenter CreateUserPresenter,
	t time.Duration,
) CreateUserUsecase {
	return CreateUserInteractor{
		repo: repo,
		ctxTimeout: t,
	}
}

func (u CreateUserInteractor) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	exsistingUser, err := u.repo.FindByEmail(ctx, input.Email)
	if exsistingUser != nil {
		return u.presenter.Output(domain.User{}), err
	}

	var user = domain.NewUser(
		primitive.NewObjectID(),
		input.Name,
		input.Email,
		input.Password,
	)
	err = u.repo.Create(ctx, user)
	if err != nil {
		return u.presenter.Output(domain.User{}), err
	}

	return CreateUserOutput{
		ID: user.ID.Hex(),
		Name: user.Name,
		Email: user.Email,
	}, nil
}