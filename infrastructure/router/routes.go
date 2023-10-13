package router

import (
	"nu/corpus-reader/adapter/api/action"

	"github.com/gin-gonic/gin"
)

func (g webEngine) healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.Healthcheck(c.Writer, c.Request)
	}
}


func (g webEngine) search() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
