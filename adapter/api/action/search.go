package action

import (
	"net/http"
)


type PatternSearchAction struct {
}


func NewPatternSearchAction() PatternSearchAction {
  return PatternSearchAction{
  }
}

func (p PatternSearchAction) PatternSearch(w http.ResponseWriter, r *http.Request) {
}
