package main

import (
	"net/http"

	"github.com/effortless-technologies/elt-delivery-api/server"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Tyler!\n")
	})

	e.GET("/delivery-quote", server.GetDeliveryQuote)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
