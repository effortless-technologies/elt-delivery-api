package main

import (
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
	e.GET("/deliveries/quote", server.GetDeliveryQuote)
	e.POST("/deliveries/:quote_id", server.CreateDelivery)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
