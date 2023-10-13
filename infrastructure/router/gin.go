package router

import (
	"context"
	"fmt"
	"net/http"
	"nu/corpus-reader/adapter/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type webEngine struct {
	router     *gin.Engine
  log         logger.Logger
	port       Port
	ctxTimeout time.Duration
}

func newServer(
	port Port,
	t time.Duration,
  log logger.Logger,
) *webEngine {
	return &webEngine{
		router:     gin.New(),
    log: log,
		port:       port,
		ctxTimeout: t,
	}
}

func (g webEngine) Listen() {
	gin.SetMode(gin.DebugMode)
	gin.Recovery()

	g.setAppHandlers(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
   	g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
      g.log.WithError(err).Fatalln("Error starting HTTP server")}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
    g.log.WithError(err).Fatalln("Error starting HTTP server")}
	}


func (g webEngine) setAppHandlers(router *gin.Engine) {
	router.GET("v1/healthcheck/", g.healthcheck())
}
