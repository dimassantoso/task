package app

import (
	"database/sql"

	movieHandlerPackage "microservice-assignment/src/movie/v1/handler"
	movieRepositoryPackage "microservice-assignment/src/movie/v1/repository"
	movieUseCasePackage "microservice-assignment/src/movie/v1/usecase"
	"sync"
)

type App struct {
	MovieHandler *movieHandlerPackage.HTTPMovieHandler
	MovieUseCase movieUseCasePackage.MovieUseCase
}

func MakeHandler(readDB *sql.DB, writeDB *sql.DB) *App {
	movieNewRepository := movieRepositoryPackage.NewMovieRepositoryMySQL(readDB, writeDB)
	movieNewUseCase := movieUseCasePackage.NewMovieUseCase(movieNewRepository)
	movieNewHandler := movieHandlerPackage.NewHTTPHandler(movieNewUseCase)

	return &App{
		MovieHandler: movieNewHandler,
		MovieUseCase: movieNewUseCase,
	}
}

func (a *App) Start() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		a.HTTPServerMain()
	}()

	wg.Wait()
}
