package usecase

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"microservice-assignment/sdk"
	sdkMocks "microservice-assignment/sdk/mocks"
	"microservice-assignment/shared"
	"strconv"
	"testing"
)

type testCase struct {
	name          string
	wantSDKResult SDKResult
	request       interface{}
	isResultError bool
}

type SDKResult struct {
	GetDetail, GetAll shared.CommonResult
}

func TestNewMovieUseCase(t *testing.T) {
	_ = NewMovieUseCase(nil)
	assert.NoError(t, nil)
}

func TestMovieUsecaseImpl_GetDetailUseCase(t *testing.T) {
	positiveSDKResult := SDKResult{GetDetail: shared.CommonResult{
		Result: sdk.ResponseDetail{
			Detail: sdk.Detail{
				Data: sdk.Data{
					IMDBID: "",
					Title:  "The 7th Rule: Captain Nog Forever",
					Year:   "2019–",
					Type:   "series",
					Poster: "https://m.media-amazon.com/images/M/MV5BYjY5MzhkODQtMDBiZi00NjhlLTk0ODYtZGE0MTM0ZGQxY2Q5XkEyXkFqcGdeQXVyMTQ4NDgxMzM@._V1_SX300.jpg",
				},
				Rated:    "N/A",
				Released: "11 Oct 2019",
				Runtime:  "N/A",
				Genre:    "Sci-Fi",
				Director: "N/A",
				Writer:   "N/A",
				Actors:   "Ryan T. Husk, Cirroc Lofton",
				Plot:     "N/A",
				Language: "English",
				Country:  "USA",
				Awards:   "N/A",
				Ratings: &[]sdk.Rating{
					{
						Source: "Internet Movie Database",
						Value:  "8.7/10",
					},
				},
				Metascore:  "N/A",
				IMDBRating: "8.7",
				IMDBVotes:  "31",
			},
		}},
	}
	negativeSDKResult := SDKResult{GetDetail: shared.CommonResult{Error: fmt.Errorf("Error")}}
	failedParseSDKResult := SDKResult{GetDetail: shared.CommonResult{Result: nil}}

	tests := []testCase{
		{
			name:          "TestCase #%s: Success",
			wantSDKResult: positiveSDKResult,
			isResultError: false,
		},
		{
			name:          "TestCase #%s: Failed (Error Result SDK)",
			wantSDKResult: negativeSDKResult,
			isResultError: true,
		},
		{
			name:          "TestCase #%s: Failed (Failed Parse)",
			wantSDKResult: failedParseSDKResult,
			isResultError: true,
		},
	}

	for id, tt := range tests {
		testCaseID := strconv.Itoa(id + 1)
		testCaseName := fmt.Sprintf(tt.name, testCaseID)
		t.Run(testCaseName, func(t *testing.T) {
			mockSDK := new(sdkMocks.OMDB)
			mockSDK.On("GetDetail", mock.Anything).Once().Return(shared.GenerateResult(tt.wantSDKResult.GetDetail.Result, tt.wantSDKResult.GetDetail.Error))

			UseCase := new(MovieUsecaseImpl)
			UseCase.OMDBService = mockSDK

			result := <-UseCase.GetDetailUseCase("tt11383906")
			shared.GetError(t, tt.isResultError, result.Error)
		})
	}
}

func TestMovieUsecaseImpl_GetAllUseCase(t *testing.T) {
	positiveSDKResult := SDKResult{GetAll: shared.CommonResult{
		Result: sdk.ResponseList{Search: []sdk.Data{{}, {}}}},
	}
	negativeSDKResult := SDKResult{GetAll: shared.CommonResult{Error: fmt.Errorf("Error")}}
	failedParseSDKResult := SDKResult{GetAll: shared.CommonResult{Result: nil}}

	tests := []testCase{
		{
			name:          "TestCase #%s: Success",
			wantSDKResult: positiveSDKResult,
			isResultError: false,
		},
		{
			name:          "TestCase #%s: Failed (Error Result SDK)",
			wantSDKResult: negativeSDKResult,
			isResultError: true,
		},
		{
			name:          "TestCase #%s: Failed (Failed Parse)",
			wantSDKResult: failedParseSDKResult,
			isResultError: true,
		},
	}

	for id, tt := range tests {
		testCaseID := strconv.Itoa(id + 1)
		testCaseName := fmt.Sprintf(tt.name, testCaseID)
		t.Run(testCaseName, func(t *testing.T) {
			mockSDK := new(sdkMocks.OMDB)
			mockSDK.On("GetAll", mock.Anything).Once().Return(shared.GenerateResult(tt.wantSDKResult.GetAll.Result, tt.wantSDKResult.GetAll.Error))

			UseCase := new(MovieUsecaseImpl)
			UseCase.OMDBService = mockSDK

			param := sdk.Param{Search: "Captain America"}
			result := <-UseCase.GetAllUseCase(param)
			shared.GetError(t, tt.isResultError, result.Error)
		})
	}
}
