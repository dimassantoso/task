package usecase

import (
	"microservice-assignment/sdk"
	"microservice-assignment/shared"
	"microservice-assignment/src/movie/v1/repository"
	"os"
)

type MovieUsecaseImpl struct {
	OMDBService    sdk.OMDB
}

type MovieUseCase interface {
	GetDetailUseCase(movieID string) <-chan shared.CommonResult
	GetAllUseCase(param sdk.Param) <-chan shared.CommonResult
}

func NewMovieUseCase(movieRepoPostgreSQL repository.MovieRepository) MovieUseCase {
	omdbBaseURL := os.Getenv("OMDB_BASE_URL")
	omdbAPIKey := os.Getenv("OMDB_API_KEY")
	omdbService := sdk.NewOMDBService(omdbBaseURL, omdbAPIKey, movieRepoPostgreSQL)
	return &MovieUsecaseImpl{
		OMDBService:    omdbService,
	}
}
