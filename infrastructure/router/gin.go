package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type webEngine struct {
	router     *gin.Engine
	port       Port
	ctxTimeout time.Duration
}

func newServer(
	port Port,
	t time.Duration,
) *webEngine {
	return &webEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: t,
	}
}

func (g webEngine) Listen() {
	fmt.Println("here")
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
		fmt.Println("Starting HTTP server")
		if err := server.ListenAndServe(); err != nil {
			fmt.Print(err)
			fmt.Println("Error while starting http server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Failed")
	}

}

func (g webEngine) setAppHandlers(router *gin.Engine) {
	router.GET("v1/healthcheck/", g.healthcheck())
}
