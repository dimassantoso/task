package model

import (
	"microservice-assignment/shared"
	"time"
)

type (
	BaseData struct {
		IMDBID string `json:"imdbId"`
		Title  string `json:"title"`
		Year   string `json:"year"`
		Type   string `json:"type"`
		Poster string `json:"poster"`
	}

	Detail struct {
		BaseData
		Rated      string    `json:"rated"`
		Released   string    `json:"released"`
		Runtime    string    `json:"runtime"`
		Genre      string    `json:"genre"`
		Director   string    `json:"director"`
		Writer     string    `json:"writer"`
		Actors     string    `json:"actors"`
		Plot       string    `json:"plot"`
		Language   string    `json:"language"`
		Country    string    `json:"country"`
		Awards     string    `json:"awards"`
		Ratings    *[]Rating `json:"ratings"`
		Metascore  string    `json:"metascore"`
		IMDBRating string    `json:"imdbRating"`
		IMDBVotes  string    `json:"imdbVotes"`
	}

	Rating struct {
		Source string `json:"source"`
		Value  string `json:"value"`
	}

	ResponseList struct {
		shared.ResponseList
		Data []BaseData `json:"data"`
	}

	LogsRequest struct {
		ID            int       `json:"id"`
		Uri           string    `json:"uri"`
		Response      string    `json:"response"`
		Created       time.Time `json:"-"`
		CreatedString string    `json:"created"`
	}
)
