package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	log.Println("Server started on :8080")
	e.Start(":8080")
}
