package infrastructure

import (
	"fmt"
	"nu/corpus-reader/infrastructure/router"
	"strconv"
	"time"
)

type config struct {
	appName       string
	ctxTimeout    time.Duration
	webServer     router.Server
	webServerPort router.Port
}

func NewConfig() *config {
	return &config{}
}

func (c *config) WebServer(instance int) *config {
	s, err := router.WebServerFactory(
		instance,
		c.webServerPort,
		c.ctxTimeout,
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
	fmt.Println("Sucessfully configured http server.")
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
