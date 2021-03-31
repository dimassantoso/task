package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"microservice-assignment/sdk"
	"microservice-assignment/shared"
	"microservice-assignment/src/movie/v1/model"
	"microservice-assignment/src/movie/v1/usecase/mocks"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type testCase struct {
	name              string
	param             sdk.Param
	wantUseCaseResult shared.CommonResult
	wantStatusCode    int
	isResultError     bool
}

func TestNewHTTPHandler(t *testing.T) {
	_ = NewHTTPHandler(new(mocks.MovieUseCase))
	assert.NoError(t, nil)
}

func TestHTTPMovieHandler_Mount(t *testing.T) {
	e := echo.New()
	handler := NewHTTPHandler(new(mocks.MovieUseCase))
	handler.Mount(e.Group("api"))

}

func TestHTTPMovieHandler_GetAllMovieHandler(t *testing.T) {
	shared.Load(shared.DirectorySchemaTest)

	var positiveUseCaseResponse = shared.CommonResult{Result: model.ResponseList{}}
	var negativeUseCaseResponse = shared.CommonResult{Error: fmt.Errorf("Error")}
	var failedParseUseCaseResponse = shared.CommonResult{Result: nil}

	positiveParam := sdk.Param{
		Page:   "2",
		Year:   "2021",
		Search: "Spider-man",
		Type:   "movie",
	}

	negativeParam := sdk.Param{
		Page:   "xyz",
		Year:   "10",
		Search: "Spider-man",
		Type:   "abc",
	}

	tests := []testCase{
		{
			name:              "TestCase #%s: Success",
			param:             positiveParam,
			wantStatusCode:    http.StatusOK,
			wantUseCaseResult: positiveUseCaseResponse,
		},
		{
			name:              "TestCase #%s: Failed (Error Result)",
			param:             positiveParam,
			wantStatusCode:    http.StatusBadRequest,
			wantUseCaseResult: negativeUseCaseResponse,
		},
		{
			name:              "TestCase #%s: Failed (Error Param Validation)",
			param:             negativeParam,
			wantStatusCode:    http.StatusBadRequest,
			wantUseCaseResult: positiveUseCaseResponse,
		},
		{
			name:              "TestCase #%s: Failed (Failed Parse Result)",
			param:             positiveParam,
			wantStatusCode:    http.StatusBadRequest,
			wantUseCaseResult: failedParseUseCaseResponse,
		},
	}

	for id, tt := range tests {
		testCaseID := strconv.Itoa(id + 1)
		testCaseName := fmt.Sprintf(tt.name, testCaseID)
		t.Run(testCaseName, func(t *testing.T) {
			mockMovieUsecase := new(mocks.MovieUseCase)
			mockMovieUsecase.On("GetAllUseCase", mock.Anything).Return(shared.GenerateResult(tt.wantUseCaseResult.Result, tt.wantUseCaseResult.Error))

			e := echo.New()
			uri := fmt.Sprintf("/api/v1/movie?search=%s&page=%s&year=%s&type=%s", tt.param.Search, tt.param.Page, tt.param.Year, tt.param.Type)
			req := httptest.NewRequest(echo.GET, uri, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := NewHTTPHandler(mockMovieUsecase)
			_ = handler.GetAllMovieHandler(c)

			assert.Equal(t, tt.wantStatusCode, rec.Code)
		})
	}
}

func TestHTTPMovieHandler_GetDetailMovieHandler(t *testing.T) {
	var positiveUseCaseResponse = shared.CommonResult{Result: model.Detail{}}
	var negativeUseCaseResponse = shared.CommonResult{Error: fmt.Errorf("Error")}
	var failedParseUseCaseResponse = shared.CommonResult{Result: nil}

	tests := []testCase{
		{
			name:              "TestCase #%s: Success",
			wantStatusCode:    http.StatusOK,
			wantUseCaseResult: positiveUseCaseResponse,
		},
		{
			name:              "TestCase #%s: Failed (Error Result)",
			wantStatusCode:    http.StatusBadRequest,
			wantUseCaseResult: negativeUseCaseResponse,
		},
		{
			name:              "TestCase #%s: Failed (Failed Parse Result)",
			wantStatusCode:    http.StatusBadRequest,
			wantUseCaseResult: failedParseUseCaseResponse,
		},
	}

	for id, tt := range tests {
		testCaseID := strconv.Itoa(id + 1)
		testCaseName := fmt.Sprintf(tt.name, testCaseID)
		t.Run(testCaseName, func(t *testing.T) {
			mockMovieUsecase := new(mocks.MovieUseCase)
			mockMovieUsecase.On("GetDetailUseCase", mock.Anything).Return(shared.GenerateResult(tt.wantUseCaseResult.Result, tt.wantUseCaseResult.Error))

			e := echo.New()
			uri := fmt.Sprintf("/api/v1/movie/%s", testCaseID)
			req := httptest.NewRequest(echo.GET, uri, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := NewHTTPHandler(mockMovieUsecase)
			_ = handler.GetDetailMovieHandler(c)

			assert.Equal(t, tt.wantStatusCode, rec.Code)
		})
	}
}
