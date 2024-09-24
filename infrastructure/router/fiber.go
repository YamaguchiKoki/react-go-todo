package router

import (
	"time"

	"github.com/YamaguchiKoki/react-go-todo/adapter/logger"
	"github.com/YamaguchiKoki/react-go-todo/adapter/presenter"
	"github.com/YamaguchiKoki/react-go-todo/adapter/repository"
	"github.com/YamaguchiKoki/react-go-todo/adapter/validator"
	"github.com/YamaguchiKoki/react-go-todo/usecase"
	"github.com/gofiber/fiber/v2"
)

type fiberEngine struct {
	app *fiber.App
	log logger.Logger
	db repository.NoSQL
	validator validator.Validator
	port Port
	ctxTimeout time.Duration
}

func newFiberServer(
	log logger.Logger,
	db repository.NoSQL,
	validator validator.Validator,
	port Port,
	t time.Duration,
) *fiberEngine {
	return &fiberEngine{
		app: fiber.New(),
		log: log,
		db: db,
		validator: validator,
		port: port,
		ctxTimeout: t,
	}
}

func (f fiberEngine) Listen() {
	f.setAppHandlers(f.app)
}

func (f fiberEngine) setAppHandlers(app *fiber.App) {
	app.Post("/v1/user", f.buildCreateUserAction)
}

//TODO:ここから
func (f fiberEngine) buildCreateUserAction(c *fiber.Ctx) fiber.H {
	uc := usecase.NewCreateUserInteractor(
		repository.NewUserNoSQL(f.db),
		presenter.NewCreateUserPresenter(),
		f.ctxTimeout,
	)
}