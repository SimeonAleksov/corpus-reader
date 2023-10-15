package repository

import (
	"testing"
)

func TestPatternSearchRepositoryImplementation_Search(t *testing.T) {
	testCases := []struct {
		name          string
		text          string
		pattern       string
		expectedCount int
	}{
		{"simple pattern", "this is johny", "john", 1},
	}

	repo := NewPatternSearchRepositoryImplementation()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := repo.Search(tc.text, tc.pattern)
			if err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if result.Count != tc.expectedCount {
				t.Errorf("Expected result count to be %d, but got %d", tc.expectedCount, result.Count)
			}
		})
	}
}
