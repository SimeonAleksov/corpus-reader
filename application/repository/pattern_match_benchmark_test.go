package repository

import (
	"testing"
)

func BenchmarkSearch(b *testing.B) {
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
		for i := 0; i < b.N; i++ {
			kmp := NewMatch(testCase.pattern)
			kmp.search(testCase.text)
		}
	}
}
