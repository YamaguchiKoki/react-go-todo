package router

import "errors"

type Server interface {
	listen()
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
	
)