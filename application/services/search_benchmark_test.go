package services

import (
	"nu/corpus-reader/application/repository"
	"testing"
)

func BenchmarkSearchInDirectory(b *testing.B) {
	testCases := []struct {
		pattern   string
		directory string
	}{
		{
			pattern:   "John",
			directory: "../../corpus",
		},
		{
			pattern:   "implementation",
			directory: "../../corpus",
		},
		{
			pattern:   "the",
			directory: "../../corpus",
		},
		{
			pattern:   "something",
			directory: "../../corpus",
		},
		{
			pattern:   "configuration",
			directory: "../../corpus",
		},
		{
			pattern:   "simeon",
			directory: "../../corpus",
		},
		{
			pattern:   "wonderware",
			directory: "../../corpus",
		},
		{
			pattern:   "Wonderware",
			directory: "../../corpus",
		},
	}

	repo := repository.NewFactory().CreateRepository(repository.KMPSearch)
	service := NewPatternSearchService(repo)
	for _, testCase := range testCases {
		for i := 0; i < b.N; i++ {
			service.SearchInDirectory(testCase.directory, testCase.pattern)
		}
	}
}
