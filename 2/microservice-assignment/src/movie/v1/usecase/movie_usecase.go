package usecase

import (
	"fmt"
	"microservice-assignment/sdk"
	"microservice-assignment/shared"
	"microservice-assignment/src/movie/v1/model"
	"strconv"
)

func (m *MovieUsecaseImpl) GetDetailUseCase(movieID string) <-chan shared.CommonResult {

	output := make(chan shared.CommonResult)
	go func() {
		defer close(output)

		getDetailResult := <-m.OMDBService.GetDetail(movieID)
		if getDetailResult.Error != nil {
			output <- shared.CommonResult{Error: getDetailResult.Error}
			return
		}

		responseDetail, ok := getDetailResult.Result.(sdk.ResponseDetail)
		if !ok {
			err := fmt.Errorf(shared.ErrorParseData)
			output <- shared.CommonResult{Error: err}
			return
		}

		detail := model.Detail{
			Rated:      responseDetail.Rated,
			Released:   responseDetail.Released,
			Runtime:    responseDetail.Runtime,
			Genre:      responseDetail.Genre,
			Director:   responseDetail.Director,
			Writer:     responseDetail.Writer,
			Actors:     responseDetail.Actors,
			Plot:       responseDetail.Plot,
			Language:   responseDetail.Language,
			Country:    responseDetail.Country,
			Awards:     responseDetail.Awards,
			Metascore:  responseDetail.Metascore,
			IMDBRating: responseDetail.IMDBRating,
			IMDBVotes:  responseDetail.IMDBVotes,
		}

		detail.IMDBID = responseDetail.IMDBID
		detail.Title = responseDetail.Title
		detail.Year = responseDetail.Year
		detail.Type = responseDetail.Type
		detail.Poster = responseDetail.Poster

		var ratings []model.Rating
		for _, value := range *responseDetail.Ratings {
			var rating model.Rating

			rating.Source = value.Source
			rating.Value = value.Value

			ratings = append(ratings, rating)
		}

		detail.Ratings = &ratings

		output <- shared.CommonResult{Result: detail}

	}()

	return output
}

func (m *MovieUsecaseImpl) GetAllUseCase(param sdk.Param) <-chan shared.CommonResult {

	output := make(chan shared.CommonResult)
	go func() {
		defer close(output)

		var (
			result        []model.BaseData
			movieResponse model.ResponseList
		)

		fetchResult := <-m.OMDBService.GetAll(param)
		if fetchResult.Error != nil {
			output <- shared.CommonResult{Error: fetchResult.Error}
			return
		}

		responseList, ok := fetchResult.Result.(sdk.ResponseList)
		if !ok {
			err := fmt.Errorf(shared.ErrorParseData)
			output <- shared.CommonResult{Error: err}
			return
		}

		movieResponse.Data = make([]model.BaseData, 0)
		for _, value := range responseList.Search {
			var movie model.BaseData

			movie.Type = value.Type
			movie.Year = value.Year
			movie.IMDBID = value.IMDBID
			movie.Poster = value.Poster
			movie.Title = value.Title

			result = append(result, movie)
		}

		total, _ := strconv.Atoi(responseList.TotalResult)
		page, _ := strconv.Atoi(param.Page)
		metaData := shared.Meta{
			TotalRecords: total,
			Page:         page,
		}

		movieResponse.Message = "Success Get Data Movie"
		movieResponse.Meta = &metaData
		movieResponse.Data = result

		output <- shared.CommonResult{Result: movieResponse}

	}()

	return output
}
