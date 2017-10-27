package server

import (
	"net/http"

	"github.com/effortless-technologies/elt-delivery-api/clients"

	"github.com/labstack/echo"
)

func GetDeliveryQuote(c echo.Context) error {

	var pickup_address, dropoff_address string
	for key, value := range c.QueryParams() {
		if key == "pickup-address" {
			pickup_address = value[0]
		} else if key == "dropoff-address" {
			dropoff_address = value[0]
		}
	}

	if pickup_address == "" {
		return c.JSON(http.StatusBadRequest, "Bad Request, need the " +
			"following query parameter: pickup_address")
	} else if dropoff_address == "" {
		return c.JSON(http.StatusBadRequest,
			"Bad Request, need the following query parameter: " +
				"dropoff_address")
	}

	response, _ := clients.GetDeliveryQuote(pickup_address, dropoff_address)

	return c.JSON(http.StatusCreated, response)
}
