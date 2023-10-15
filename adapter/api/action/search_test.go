package action

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"nu/corpus-reader/application/presenter"
	"nu/corpus-reader/application/repository"
	"nu/corpus-reader/application/usecase"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	type args struct {
		payload []byte
	}
	testCases := []struct {
		name               string
		args               args
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "Search simple pattern",
			args: args{
				payload: []byte(
					`{
						"word": "John",
						"directory": "../../../corpus"
					}`,
				),
			},
			expectedBody:       `{"count":12}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Search simple pattern",
			args: args{
				payload: []byte(
					`{
						"word": "simeon",
						"directory": "../../../corpus"
					}`,
				),
			},
			expectedBody:       `{"count":0}`,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req, _ := http.NewRequest(
				http.MethodPost,
				"/counter",
				bytes.NewReader(testCase.args.payload),
			)

			var (
				rr      = httptest.NewRecorder()
				handler = http.NewServeMux()
			)

			repo := repository.NewFactory().CreateRepository(repository.KMPSearch)
			uc := usecase.NewCreatePatternSearchInteractor(
				repo,
				presenter.NewCreatePatternSearchPresenter(),
				0,
			)

			ac := NewPatternSearchAction(uc)

			handler.HandleFunc("/counter", ac.PatternSearch)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("The handler returned an unexpected HTTP status code: returned '%v' expected '%v'",
					status,
					http.StatusOK,
				)
			}

			var result = strings.TrimSpace(rr.Body.String())
			if !strings.EqualFold(result, testCase.expectedBody) {
				t.Errorf(
					"[TestCase '%s'] Result: '%v' | Expected: '%v'",
					testCase.name,
					result,
					testCase.expectedBody,
				)
			}
		},
		)
	}
}
