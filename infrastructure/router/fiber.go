package router

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YamaguchiKoki/react-go-todo/adapter/api/action"
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

//FIXME
func (f fiberEngine) Listen() {
	f.setAppHandlers(f.app)


	// シグナルチャンネルの作成
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Fiber サーバーを起動
	go func() {
		f.log.WithFields(logger.Fields{"port": f.port}).Infof("Starting HTTP Server")
		if err := f.app.Listen(fmt.Sprintf(":%d", f.port)); err != nil {
			f.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	// シャットダウンシグナルを待機
	<-stop

	// Graceful Shutdown のためのタイムアウト付きコンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := f.app.Shutdown(); err != nil {
		f.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	f.log.Infof("Service down")
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

	act := action.NewCreateUserAction(uc, f.log, f.validator)

	act.Execute()
}