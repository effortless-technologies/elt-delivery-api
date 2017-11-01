package main

import (
	"github.com/effortless-technologies/elt-delivery-api/server"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/deliveries/quote", server.GetDeliveryQuote)
	e.POST("/deliveries/:quote_id", server.CreateDelivery)
	e.GET("/deliveries/:delivery_id", server.GetDelivery)
	e.GET("/deliveries", server.GetDeliveries)

	e.Logger.Fatal(e.Start(":1323"))
}
