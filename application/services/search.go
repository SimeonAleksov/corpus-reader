package services

import (
	"nu/corpus-reader/application/domain"
	"nu/corpus-reader/application/repository"
	"nu/corpus-reader/infrastructure/log"
)


var logger = log.NewLogrusLogger("PATTERN-SEARCH")

type PatternSearchService struct {
  patternSearchRepo repository.PatternSearchRepository
}


func NewPatternSearchService(repo repository.PatternSearchRepository) *PatternSearchService {
  return &PatternSearchService{
    patternSearchRepo: repo,
  }
}

func (s *PatternSearchService) Search(text string, pattern string) (*domain.PatternSearchResult, error) {
  count, err := s.patternSearchRepo.Search(text, pattern)
  if err != nil {
      logger.WithError(err).Fatalln("Error while searching pattern in text")
  }

  return count, nil
}
