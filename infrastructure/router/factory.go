package router

import (
	"errors"
	"nu/corpus-reader/adapter/logger"
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
	log logger.Logger,
) (Server, error) {
	switch instance {
	case Gin:
		return newServer(port, ctxTimeout, log), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
