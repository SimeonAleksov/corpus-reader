package action

import (
	"encoding/json"
	"net/http"
	"nu/corpus-reader/adapter/api/response"
	"nu/corpus-reader/application/usecase"
)

type PatternSearchAction struct {
	uc usecase.PatternSearchUseCase
}

func NewPatternSearchAction(
	uc usecase.PatternSearchUseCase,
) PatternSearchAction {
	return PatternSearchAction{
		uc: uc,
	}
}

func (p PatternSearchAction) PatternSearch(w http.ResponseWriter, r *http.Request) {
	var input usecase.PatternSearchInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	if err := input.Validate(); err != nil {
		response.NewError(err.Error, err.Status).Send(w)
		return
	}

	defer r.Body.Close()

	output, err := p.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
