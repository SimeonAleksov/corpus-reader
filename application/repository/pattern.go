package repository

import (
	"nu/corpus-reader/application/domain"
)

type PatternSearchRepository interface {
  Search(Text string, Pattern string) (*domain.PatternSearchResult, error)
}

type PatternSearchRepositoryImplementation struct {

}

func NewPatternSearchRepositoryImplementation() *PatternSearchRepositoryImplementation {
  return &PatternSearchRepositoryImplementation{}
}

func (repo *PatternSearchRepositoryImplementation) Search(text string, pattern string) (*domain.PatternSearchResult, error) {
  patternMatch := NewMatch(pattern)
  count := patternMatch.search(text)
  result := domain.NewPatternSearchResult(count)

  return &result, nil
}
