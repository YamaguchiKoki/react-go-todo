package infrastructure

import (
	"time"

	"github.com/YamaguchiKoki/react-go-todo/adapter/logger"
	"github.com/YamaguchiKoki/react-go-todo/adapter/repository"
	"github.com/YamaguchiKoki/react-go-todo/adapter/validator"
	"github.com/YamaguchiKoki/react-go-todo/infrastructure/router"
)

type config struct {
	appName string
	logger logger.Logger
	validator validator.Validator
	dbSQL repository.SQL
	dbNoSQL repository.NoSQL
	ctxTimeout time.Duration
	webServerPort router.Port
}