package presenter

import (
	"nu/corpus-reader/application/domain"
	"nu/corpus-reader/application/usecase"
)


type createPatternSearchPresenter struct {}

func NewCreatePatternSearchPresenter() createPatternSearchPresenter {
  return createPatternSearchPresenter{}
}

func (p createPatternSearchPresenter) Output(result domain.PatternSearchResult) usecase.PatternSearchOutput {
  return usecase.PatternSearchOutput{
    Count: result.Count,
  }
}
