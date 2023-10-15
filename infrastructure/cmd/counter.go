package cmd

import (
	"nu/corpus-reader/application/presenter"
	"nu/corpus-reader/application/repository"
	"nu/corpus-reader/application/usecase"
	"nu/corpus-reader/infrastructure/log"

	"github.com/spf13/cobra"
)

var word string
var directory string

var logger = log.NewLogrusLogger("COUNTER-COMMAND")

var counterCommand = &cobra.Command{
	Use:   "counter",
	Short: "Count word in directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("Searching for: %v in %v", word, directory)
		input := usecase.PatternSearchInput{
			Pattern:   word,
			Directory: directory,
		}
		repo := repository.NewFactory().CreateRepository(repository.KMPSearch)
		uc := usecase.NewCreatePatternSearchInteractor(
			repo,
			presenter.NewCreatePatternSearchPresenter(),
			0,
		)
		output, err := uc.Execute(cmd.Context(), input)
		if err != nil {
			logger.WithError(err).Fatalln("Error while searching for pattern.")
		}
		logger.Infof("count: %v", output.Count)
	},
}
