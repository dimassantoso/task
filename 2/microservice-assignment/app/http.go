package app

import (
	"fmt"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"microservice-assignment/middleware"
	"os"
	"strconv"
)

const DefaultPort = 8080

// HTTPServerMain main function for serving services over http
func (app *App) HTTPServerMain() {
	e := echo.New()
	e.Use(middleware.ServerHeader, middleware.Logger)
	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	e.Debug = true
	handlerGroup := e.Group("/api")
	handlerGroup.Use(middleware.BearerVerify())

	app.MovieHandler.Mount(handlerGroup)

	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = DefaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))
}
