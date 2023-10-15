package router

import (
	"nu/corpus-reader/adapter/api/action"
	"nu/corpus-reader/application/presenter"
	"nu/corpus-reader/application/repository"
	"nu/corpus-reader/application/usecase"

	"github.com/gin-gonic/gin"
)

func (g webEngine) healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.Healthcheck(c.Writer, c.Request)
	}
}

func (g webEngine) search() gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := repository.NewFactory().CreateRepository(repository.KMPSearch)
		uc := usecase.NewCreatePatternSearchInteractor(
			repo,
			presenter.NewCreatePatternSearchPresenter(),
			g.ctxTimeout,
		)
		action.NewPatternSearchAction(uc).PatternSearch(c.Writer, c.Request)
	}
}
