package domain

type (
  PatternSearchResult struct {
    Count int `json:"count"`
  }
)

func NewPatternSearchResult(count int) PatternSearchResult {
  return PatternSearchResult{
    Count: count,
  }
}
