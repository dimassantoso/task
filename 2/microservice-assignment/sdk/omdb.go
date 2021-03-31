package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"microservice-assignment/shared"
	"microservice-assignment/src/movie/v1/model"
	"microservice-assignment/src/movie/v1/repository"
	"net/url"
)

type OMDB interface {
	GetDetail(id string) <-chan shared.CommonResult
	GetAll(param Param) <-chan shared.CommonResult
}

type OMDBService struct {
	BaseURL *url.URL
	APIKey  string
	Logs repository.MovieRepository
}

func NewOMDBService(baseURL string, apiKey string, logRepository repository.MovieRepository) OMDB {
	service := new(OMDBService)

	service.APIKey = apiKey

	var err error
	service.BaseURL, err = url.Parse(baseURL)
	if err != nil {
		err = errors.New("url is invalid")
		panic(err)
	}

	service.Logs = logRepository

	return service
}

func (O *OMDBService) GetAll(param Param) <-chan shared.CommonResult {
	var (
		response ResponseList
		result   shared.CommonResult
	)

	output := make(chan shared.CommonResult)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				result.Error = fmt.Errorf("%v", r)
			}
			output <- result
			close(output)
		}()

		uri := fmt.Sprintf("%s?apikey=%s&s=%s&p=%s&y=%s&type=%s", O.BaseURL, O.APIKey, param.Search, param.Page, param.Year, param.Type)

		r, _, err := shared.NewHTTPRequest().Do(context.Background(), "GET", uri, nil, nil)
		if err != nil {
			panic(err)
		}

		_ = json.Unmarshal(r, &response)
		if response.Response != "True" {
			panic(response.Error)
		}

		result.Result = response

		responseJSON, _ := json.Marshal(response)
		logRequest := model.LogsRequest{
			Uri: uri,
			Response: string(responseJSON),
		}

		O.Logs.SaveLog(&logRequest)

	}()

	return output

}

func (O *OMDBService) GetDetail(id string) <-chan shared.CommonResult {
	var (
		response ResponseDetail
		result   shared.CommonResult
	)

	output := make(chan shared.CommonResult)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				result.Error = fmt.Errorf("%v", r)
			}
			output <- result
			close(output)
		}()

		uri := fmt.Sprintf("%s?apiKey=%s&i=%s", O.BaseURL, O.APIKey, id)

		r, _, err := shared.NewHTTPRequest().Do(context.Background(), "GET", uri, nil, nil)
		if err != nil {
			panic(err)
		}

		_ = json.Unmarshal(r, &response)
		if response.Response != "True" {
			panic(response.Error)
		}

		result.Result = response

		responseJSON, _ := json.Marshal(response)
		logRequest := model.LogsRequest{
			Uri: uri,
			Response: string(responseJSON),
		}

		O.Logs.SaveLog(&logRequest)
	}()

	return output
}
