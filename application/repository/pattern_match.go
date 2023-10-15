package repository

type Match struct {
	pattern       string
	patternLength int
	partialMatch  []int
}

// Build a new pattern table for KMP searching
func NewMatch(pattern string) *Match {
	return &Match{
		pattern:       pattern,
		patternLength: len(pattern),
		partialMatch:  buildPartialMatchTable(pattern),
	}
}

func buildPartialMatchTable(pattern string) []int {
	word_length := len(pattern)
	partialMatch := make([]int, word_length)
	j := 0

	for i := 1; i < word_length; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = partialMatch[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}
		partialMatch[i] = j
	}

	return partialMatch
}

func (m *Match) search(text string) int {
	textLength := len(text)

	if m.patternLength == 0 {
		return 0
	}

	if textLength == 0 {
		return 0
	}

	count := 0
	textIndex := 0
	patternIndex := 0

	for textIndex < textLength {
		for patternIndex > 0 && text[textIndex] != m.pattern[patternIndex] {
			patternIndex = m.partialMatch[patternIndex-1]
		}
		if text[textIndex] == m.pattern[patternIndex] {
			if patternIndex == m.patternLength-1 {
				count++
				patternIndex = m.partialMatch[patternIndex]
			} else {
				patternIndex++
			}
		}
		textIndex++
	}

	return count
}
