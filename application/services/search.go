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
		return nil, err
	}

	return count, nil
}

func (s *PatternSearchService) SearchInDirectory(directory string, pattern string) (*domain.PatternSearchResult, error) {
	dirRepo := repository.NewDirectoryRepository()
	res, err := dirRepo.ListFiles(directory, []string{"txt"})
	if err != nil {
		logger.WithError(err).Fatalln("Error while listing files.")
	}
	count := 0
	for _, filepath := range res.Files {
		content, err := dirRepo.GetFileContent(filepath)
		if err != nil {
			logger.WithError(err).Fatalln("Error while reading content.")
		}
		res, err := s.Search(content, pattern)
		if err != nil {
			logger.WithError(err).Fatalln("Error while searching for pattern.")
		}
		count = count + res.Count
	}
	return &domain.PatternSearchResult{
		Count: count,
	}, nil
}
