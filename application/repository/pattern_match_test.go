package repository

import (
	"fmt"
	"testing"
)

func TestBuildPartialMatchTable(t *testing.T) {
	testCases := []struct {
		pattern  string
		expected []int
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
		t.Run(testCase.pattern, func(t *testing.T) {
			actual := buildPartialMatchTable(testCase.pattern)
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

func TestSearch(t *testing.T) {
	testCases := []struct {
		text, pattern string
		expected      int
	}{
		{"ABABDABACDABABCABAB", "ABABCABAB", 1},
		{"AABAACAADAABAAABAA", "AABA", 3},
		{"ABCABCABCABC", "ABCA", 3},
		{"AAAAA", "AA", 4},
		{"ABCDE", "XYZ", 0},
		{"AAAA", "A", 4},
		{"ABABABA", "ABA", 3},
		{"ABCDEF", "", 0},
		{"this is john and i am johnys friend, we are the best johns", "john", 3},
	}

	for _, testCase := range testCases {
		kmp := NewMatch(testCase.pattern)
		t.Run(fmt.Sprintf("Text: %s, pattern: %s", testCase.text, testCase.pattern), func(t *testing.T) {
			actual := kmp.search(testCase.text)
			if actual != testCase.expected {
				t.Errorf("Expected %d occurrences but got %d", testCase.expected, actual)
			}
		})
	}
}
