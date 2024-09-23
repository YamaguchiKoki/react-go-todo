package router

import (
	"errors"
	"time"

	"github.com/YamaguchiKoki/react-go-todo/adapter/logger"
	"github.com/YamaguchiKoki/react-go-todo/adapter/repository"
	"github.com/YamaguchiKoki/react-go-todo/adapter/validator"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGorillaMux int = iota
	InstanceFiber
)

func NewWebServerFactory(
	instance int,
	log logger.Logger,
	dbSQL repository.SQL,
	dbNoSQL repository.NoSQL,
	validator validator.Validator,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	// case InstanceGorillaMux:
	// 	return newGorillaMux(log, dbSQL, validator, port, ctxTimeout)
	case InstanceFiber:
		return newFiber(log, dbNoSQL, validator, port, ctxTimeout)
	}
}