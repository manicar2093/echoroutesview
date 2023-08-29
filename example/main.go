package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manicar2093/echoroutesview"
)

func main() {
	e := echo.New()

	e.GET("/one", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "/one")
	})

	e.POST("/two", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "/two")
	})

	e.PUT("/three", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "/three")
	})

	e.DELETE("/four", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "/four")
	})

	println(e.Routes())
	if err := echoroutesview.RegisterRoutesViewer(e); err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":7777"))
}
