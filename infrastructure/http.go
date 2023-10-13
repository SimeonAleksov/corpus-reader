package infrastructure

import (
	"nu/corpus-reader/adapter/logger"
	"nu/corpus-reader/infrastructure/log"
	"nu/corpus-reader/infrastructure/router"
	"strconv"
	"time"
)

type config struct {
	appName       string
  logger        logger.Logger
	ctxTimeout    time.Duration
	webServer     router.Server
	webServerPort router.Port
}

func NewConfig() *config {
	return &config{}
}


func (c *config) Logger() *config {
  log := log.NewLogrusLogger("CONFIG")
  c.logger = log
  c.logger.Infof("Configured logrus logger.")
  return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.WebServerFactory(
		instance,
		c.webServerPort,
		c.ctxTimeout,
    c.logger,
	)
	if err != nil {
	}
	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
	}

	c.webServerPort = router.Port(p)
  c.logger.Infof("Sucessfully configured http server.")
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
