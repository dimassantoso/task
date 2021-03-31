package sdk

type (
	Data struct {
		IMDBID string `json:"imdbID,omitempty"`
		Title  string `json:"Title,omitempty"`
		Year   string `json:"Year,omitempty"`
		Type   string `json:"Type,omitempty"`
		Poster string `json:"Poster,omitempty"`
	}

	Detail struct {
		Data
		Rated      string    `json:"Rated,omitempty"`
		Released   string    `json:"Released,omitempty"`
		Runtime    string    `json:"Runtime,omitempty"`
		Genre      string    `json:"Genre,omitempty"`
		Director   string    `json:"Director,omitempty"`
		Writer     string    `json:"Writer,omitempty"`
		Actors     string    `json:"Actors,omitempty"`
		Plot       string    `json:"Plot,omitempty"`
		Language   string    `json:"Language,omitempty"`
		Country    string    `json:"Country,omitempty"`
		Awards     string    `json:"Awards,omitempty"`
		Ratings    *[]Rating `json:"Ratings,omitempty"`
		Metascore  string    `json:"Metascore,omitempty"`
		IMDBRating string    `json:"imdbRating,omitempty"`
		IMDBVotes  string    `json:"imdbVotes,omitempty"`
	}

	Rating struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	}

	BaseResponse struct {
		Response string `json:"Response"`
		Error    string `json:"Error,omitempty"`
	}

	ResponseDetail struct {
		BaseResponse
		Detail
	}

	ResponseList struct {
		BaseResponse
		TotalResult string `json:"totalResults"`
		Search      []Data `json:"Search"`
	}

	Param struct {
		Page   string `json:"page" default:"1"`
		Search string `json:"search"`
		Type   string `json:"type"`
		Year   string `json:"year"`
	}
)
