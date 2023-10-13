package repository

import (
	"testing"
)

func TestBuildPartialMatchTable(t *testing.T) {
	testCases := []struct {
		word         string
		expected     []int
	}{
		{"abcabc", []int{0, 0, 0, 1, 2, 3}},
		{"abcdabca", []int{0, 0, 0, 0, 1, 2, 3, 1}},
		{"aaaaa", []int{0, 1, 2, 3, 4}},
		{"abababab", []int{0, 0, 1, 2, 3, 4, 5, 6}},
		{"john", []int{0, 0, 0, 0}},
		{"acacagt", []int{0, 0, 1, 2, 3, 0, 0}},
    {"abracadabra", []int{0, 0, 0, 1, 0, 1, 0, 1, 2, 3, 4}},
    {"banana", []int{0, 0, 0, 0, 0, 0}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.word, func(t *testing.T) {
			actual := buildPartialMatchTable(testCase.word)
			if len(actual) != len(testCase.expected) {
				t.Errorf("Expected partial match table length %d, but got %d", len(testCase.expected), len(actual))
			}
			for i, val := range actual {
				if val != testCase.expected[i] {
					t.Errorf("Expected partial match table at index %d to be %d, but got %d", i, testCase.expected[i], val)
				}
			}
		})
	}
}
