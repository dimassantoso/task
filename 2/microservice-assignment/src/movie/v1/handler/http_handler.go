package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"microservice-assignment/sdk"
	"microservice-assignment/shared"
	"microservice-assignment/src/movie/v1/model"
	"microservice-assignment/src/movie/v1/usecase"
	"net/http"
)

type HTTPMovieHandler struct {
	MovieUseCase usecase.MovieUseCase
}

func NewHTTPHandler(movieUsecase usecase.MovieUseCase) *HTTPMovieHandler {
	return &HTTPMovieHandler{
		MovieUseCase: movieUsecase,
	}
}

func (h *HTTPMovieHandler) Mount(group *echo.Group) {
	group.GET("/v1/movie", h.GetAllMovieHandler)
	group.GET("/v1/movie/:movieID", h.GetDetailMovieHandler)
}

func (h *HTTPMovieHandler) GetAllMovieHandler(c echo.Context) error {
	var params sdk.Param

	_ = shared.ParseFromQueryParam(c.Request().URL.Query(), &params)
	err := shared.Validate("movie_list_params", params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseDetailOutput(false, http.StatusBadRequest, err.Error(), nil))
	}

	result := <-h.MovieUseCase.GetAllUseCase(params)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseDetailOutput(false, http.StatusBadRequest, result.Error.Error(), nil))
	}

	response, ok := result.Result.(model.ResponseList)
	if !ok {
		return c.JSON(http.StatusBadRequest, shared.ResponseDetailOutput(false, http.StatusBadRequest, shared.ErrorParseData, nil))
	}
	response.Success = true
	response.Code = http.StatusOK

	return c.JSON(http.StatusOK, response)

}

func (h *HTTPMovieHandler) GetDetailMovieHandler(c echo.Context) error {
	result := <-h.MovieUseCase.GetDetailUseCase(c.Param("movieID"))
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, shared.ResponseDetailOutput(false, http.StatusBadRequest, result.Error.Error(), nil))
	}

	response, ok := result.Result.(model.Detail)
	if !ok {
		err := fmt.Errorf(shared.ErrorParseData)
		return c.JSON(http.StatusBadRequest, shared.ResponseDetailOutput(false, http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, shared.ResponseDetailOutput(true, http.StatusCreated, "Success Get Data Detail Movie", response))
}
