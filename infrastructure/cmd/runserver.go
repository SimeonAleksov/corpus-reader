package cmd

import (
	"nu/corpus-reader/infrastructure"
	"nu/corpus-reader/infrastructure/router"

	"github.com/spf13/cobra"
)

var runserverCommand = &cobra.Command{
	Use:   "runserver",
	Short: "Start a http server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var app = infrastructure.NewConfig().Logger()
		app.WebServerPort("8080").WebServer(router.Gin).Start()
	},
}
