package usecase

import (
	"context"
	"nu/corpus-reader/application/domain"
	"nu/corpus-reader/application/repository"
	"nu/corpus-reader/application/services"
	"time"
)

type (
	PatternSearchUseCase interface {
		Execute(context.Context, PatternSearchInput) (PatternSearchOutput, error)
	}
	PatternSearchInput struct {
		Pattern   string `json:"word"`
		Directory string `json:"directory"`
	}
	PatternSearchPresenter interface {
		Output(domain.PatternSearchResult) PatternSearchOutput
	}
	PatternSearchOutput struct {
		Count int `json:"count"`
	}
	createPatternSearchInteractor struct {
		repo       repository.PatternSearchRepository
		presenter  PatternSearchPresenter
		ctxTimeout time.Duration
	}
)

func NewCreatePatternSearchInteractor(
	repo repository.PatternSearchRepository,
	presenter PatternSearchPresenter,
	t time.Duration,
) createPatternSearchInteractor {
	return createPatternSearchInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (p createPatternSearchInteractor) Execute(ctx context.Context, input PatternSearchInput) (PatternSearchOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, p.ctxTimeout)
	defer cancel()
	result, err := services.NewPatternSearchService(p.repo).SearchInDirectory(
		input.Directory,
		input.Pattern,
	)
	if err != nil {
		return PatternSearchOutput{}, err
	}

	return p.presenter.Output(*result), nil
}

func (i PatternSearchInput) Validate() *domain.RestError {
	if i.Directory == "" {
		return repository.BadRequestError("Missing directory.")
	}
	return nil
}
