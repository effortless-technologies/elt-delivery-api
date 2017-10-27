package clients

import (
	"time"

	"github.com/effortless-technologies/elt-delivery-api/models"

	"github.com/dghubble/sling"
)

const PostmatesClientTimeout = 5 * time.Second
const PostmatesBaseUrl = "https://api.postmates.com"
const PostmatesAPIKey = "0c55e671-dcd6-4d95-b237-abfe27952877"

func Authenticate() error {

	// TODO: implement

	return nil
}

func GetDeliveryQuote(
	pickup_address string, dropoff_address string) (*models.Quote, error) {

	postmates_base := sling.New().Base(
		PostmatesBaseUrl).SetBasicAuth(
		PostmatesAPIKey,"").Set(
		"Accept", "application/json").Set(
		"Content-Type", "application/json")

	type form_params struct {
		PickupAddress		string		`url:"pickup_address"`
		DropoffAddress 		string 		`url:"dropoff_address"`
	}

	params := &form_params{
		PickupAddress: pickup_address,
		DropoffAddress: dropoff_address,
	}

	response := &models.Quote{}

	req, err := postmates_base.New().Post(
		"/v1/customers/cus_LSo5Dq0t8ppZFF/delivery_quotes").
		BodyForm(params).Request()
	if err != nil {
		return nil, err
	}

	_, err = postmates_base.Do(req, &response, err)
	if err != nil {
		return nil, err
	}

	return response, nil
}