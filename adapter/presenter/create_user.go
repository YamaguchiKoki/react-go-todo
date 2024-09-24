package presenter

import (
	"github.com/YamaguchiKoki/react-go-todo/domain"
	"github.com/YamaguchiKoki/react-go-todo/usecase"
)

type createUserPresenter struct{}

func NewCreateUserPresenter() usecase.CreateUserPresenter {
	return createUserPresenter{}
}

func (c createUserPresenter) Output(user domain.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
	}
}