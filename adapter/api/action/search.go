package action

import (
	"net/http"
	"nu/corpus-reader/application/repository"
	"nu/corpus-reader/application/services"
)


type PatternSearchAction struct {
}


func NewPatternSearchAction() PatternSearchAction {
  return PatternSearchAction{
  }
}

func (p PatternSearchAction) PatternSearch(w http.ResponseWriter, r *http.Request) {
  repo := repository.NewPatternSearchRepositoryImplementation()
  service := services.NewPatternSearchService(repo)
  service.Search("test", "asdad")
}
