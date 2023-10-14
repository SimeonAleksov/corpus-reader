package services

import (
	"fmt"
	"nu/corpus-reader/application/repository"
	"testing"
)


func TestSearchInDirectory(t *testing.T) {
	testCases := []struct {
		pattern string
    directory string
		expectedCount  int 
	}{
		{
			pattern: "John",
      directory: "../../corpus",
      expectedCount: 12,
		},
		{
			pattern: "implementation",
      directory: "../../corpus",
      expectedCount: 148,
		},
		{
			pattern: "the",
      directory: "../../corpus",
      expectedCount: 28656,
		},
		{
			pattern: "something",
      directory: "../../corpus",
      expectedCount: 63,
		},
		{
			pattern: "configuration",
      directory: "../../corpus",
      expectedCount: 9,
		},
		{
			pattern: "simeon",
      directory: "../../corpus",
      expectedCount: 0,
		},
		{
			pattern: "wonderware",
      directory: "../../corpus",
      expectedCount: 0,
		},
		{
			pattern: "Wonderware",
      directory: "../../corpus",
      expectedCount: 15,
		},
	}

  repo := repository.NewFactory().CreateRepository(repository.KMPSearch)
  service := NewPatternSearchService(repo)
  for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Pattern: %v", testCase.pattern), func(t *testing.T) {
			result, err := service.SearchInDirectory(testCase.directory, testCase.pattern)

			if err != nil {
				t.Errorf("Error: %v", err)
			}

      if result.Count != testCase.expectedCount {
        t.Errorf("Expected count %d, but go %d", testCase.expectedCount, result.Count)
      }
		})
  }
  
}

