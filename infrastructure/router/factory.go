package router

import (
	"errors"
	"time"
)

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

type Port int64
type Server interface {
	Listen()
}

const Gin = 1

func WebServerFactory(
	instance int,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	case Gin:
		return newServer(port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
