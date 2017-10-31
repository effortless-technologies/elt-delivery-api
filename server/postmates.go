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

func CreateDelivery(c echo.Context) error {

	// TODO: refactor lack of required body var to return 400 instead of 500

	quote_id := c.Param("quote_id")
	params := new(clients.CreateDeliveryParams)
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, params)
	}
	params.QuoteId = quote_id


	response, err := clients.CreateDelivery(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, response)
}
