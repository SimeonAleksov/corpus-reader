package repository

type RepositoryType int


const (
  KMPSearch RepositoryType = iota
)

type RepositoryFactory struct {}

func NewFactory() *RepositoryFactory {
  return &RepositoryFactory{}
}

func (f *RepositoryFactory) CreateRepository(repositoryType RepositoryType) PatternSearchRepository  {
  switch repositoryType {
  case KMPSearch:
    return NewPatternSearchRepositoryImplementation()
  default:
    return nil
  }
}
