package action

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/YamaguchiKoki/react-go-todo/adapter/api/logging"
	"github.com/YamaguchiKoki/react-go-todo/adapter/api/response"
	"github.com/YamaguchiKoki/react-go-todo/adapter/logger"
	"github.com/YamaguchiKoki/react-go-todo/adapter/validator"
	"github.com/YamaguchiKoki/react-go-todo/domain"
	"github.com/YamaguchiKoki/react-go-todo/usecase"
)

//OneActionController
type CreateUserAction struct {
	log logger.Logger
	uc usecase.CreateUserUsecase
	validator validator.Validator
	
	logKey, logMsg string
}

func NewCreateUserAction(uc usecase.CreateUserUsecase, log logger.Logger, v validator.Validator) CreateUserAction {
	return CreateUserAction{
		uc: uc,
		log: log,
		validator: v,
		logKey: "create_user",
		logMsg: "creating a new user",
	}
}

func (u CreateUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput
	//リクエストボディをinputにバインド
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			u.log,
			err,
			u.logKey,
			http.StatusBadRequest,
		).Log(u.logMsg)

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	//バリデーション
	if errs := u.validateInput(input); len(errs) > 0 {
		logging.NewError(
			u.log,
			response.ErrInvalidInput,
			u.logKey,
			http.StatusBadRequest,
		).Log(u.logMsg)

		response.NewErrorMessage(errs, http.StatusBadRequest)
		return
	}

	//ビジネスロジック実行
	output, err := u.uc.Execute(r.Context(), input)
	if err != nil {
		u.handleErr(w, err)
		return
	}

	logging.NewInfo(u.log, u.logKey, http.StatusCreated).Log(u.logMsg)

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

func (u CreateUserAction) handleErr(w http.ResponseWriter, err error) {
	switch err {
	case domain.ErrUserNotFound:
		logging.NewError(
			u.log,
			err,
			u.logKey,
			http.StatusNotFound,
		)
	default:
		logging.NewError(
			u.log,
			err,
			u.logKey,
			http.StatusInternalServerError,
		).Log(u.logMsg)

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
}

//FIXME
func (u CreateUserAction) validateInput(input usecase.CreateUserInput) []string {
	var (
		msgs              []string
		errAccountsEquals = errors.New("account origin equals destination account")
		accountIsEquals   = true
		accountsIsEmpty   = true
	)

	if !accountsIsEmpty && accountIsEquals {
		msgs = append(msgs, errAccountsEquals.Error())
	}

	err := u.validator.Validate(input)
	if err != nil {
		for _, msg := range u.validator.Messages() {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}