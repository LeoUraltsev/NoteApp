package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Docker Container!")
	})
	err := e.Start(":10000")
	if err != nil {
		return err
	}
	return nil
}
