package repository

import (
	"database/sql"
	"microservice-assignment/shared"
	"microservice-assignment/src/movie/v1/model"
)

type MovieRepoMySQL struct {
	readDB, writeDB *sql.DB
}

type MovieRepository interface {
	//Command
	SaveLog(req *model.LogsRequest) <-chan shared.CommonResult
}

func NewMovieRepositoryMySQL(readDB, writeDB *sql.DB) *MovieRepoMySQL {
	return &MovieRepoMySQL{readDB: readDB, writeDB: writeDB}
}
