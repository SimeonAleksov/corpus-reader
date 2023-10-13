package repository


type Match struct {
  partialMatch []int
}


// Build a new pattern table for KMP searching
func NewMatch(word string) *Match {
  return &Match{
    partialMatch: buildPartialMatchTable(word),
  }
}


func buildPartialMatchTable(word string) []int {
  word_length := len(word)
  partialMatch := make([]int, word_length)
  j := 0

	for i := 1; i < word_length; i++ {
		for j > 0 && word[i] != word[j] {
			j = partialMatch[j-1]
		}

		if word[i] == word[j] {
			j++
		}
		partialMatch[i] = j
	}

  return partialMatch
}
