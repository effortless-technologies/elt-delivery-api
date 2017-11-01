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

func makeBaseUrl () *sling.Sling {
	postmates_base := sling.New().Base(
		PostmatesBaseUrl).SetBasicAuth(
		PostmatesAPIKey,"").Set(
		"Accept", "application/json")

	return postmates_base
}

func GetDeliveryQuote(
	pickup_address string, dropoff_address string) (*models.Quote, error) {

	postmates_base := makeBaseUrl()
	type form_params struct {
		PickupAddress		string		`url:"pickup_address"`
		DropoffAddress 		string 		`url:"dropoff_address"`
	}

	params := &form_params{
		PickupAddress: pickup_address,
		DropoffAddress: dropoff_address,
	}

	uri := "/v1/customers/cus_LSo5Dq0t8ppZFF/delivery_quotes"

	req, err := postmates_base.New().Post(uri).BodyForm(params).Request()
	if err != nil {
		return nil, err
	}

	response := &models.Quote{}

	_, err = postmates_base.Do(req, &response, err)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type CreateDeliveryParams struct {
	QuoteId				string		`url:"quote_id,omitempty"json:"quote_id,omitempty"`
	Manifest	 		string 		`url:"manifest"json:"manifest"`
	ManifestReference	string 		`url:"manifest_reference,omitempty"json:"manifest_reference,omitempty"`
	PickupName			string		`url:"pickup_name"json:"pickup_name"`
	PickupAddress		string		`url:"pickup_address"json:"pickup_address"`
	PickupPhoneNumber	string		`url:"pickup_phone_number"json:"pickup_phone_number"`
	PickupBusinessName	string		`url:"pickup_business_name,omitempty"json:"pickup_business_name,omitempty"`
	PickupNotes			string		`url:"pickup_notes,omitempty"json:"pickup_notes,omitempty"`
	DropoffName			string		`url:"dropoff_name"json:"dropoff_name"`
	DropoffAddress		string		`url:"dropoff_address"json:"dropoff_address"`
	DropoffPhoneNumber	string		`url:"dropoff_phone_number"json:"dropoff_phone_number"`
	DropoffBusinessName	string		`url:"dropoff_business_name,omitempty"json:"dropoff_business_name,omitempty"`
	DropoffNotes		string		`url:"dropoff_notes,omitempty"json:"dropoff_notes,omitempty"`
	RequireId			string		`url:"requires_id,omitempty"json:"requires_id,omitempty"`
}

func CreateDelivery(params *CreateDeliveryParams) (*models.Delivery, error) {

	postmates_base := makeBaseUrl()
	uri := "/v1/customers/cus_LSo5Dq0t8ppZFF/deliveries"
	req, err := postmates_base.New().Post(uri).BodyForm(params).Request()
	if err != nil {
		return nil, err
	}

	response := new(models.DeliveryParse)
	_, err = postmates_base.Do(req, &response, err)
	if err != nil {
		return nil, err
	}

	d := new(models.Delivery)
	d.Parse(response)

	return d, nil
}

func GetDelivery(id string) (*models.Delivery, error) {

	postmates_base := makeBaseUrl()
	uri := "/v1/customers/cus_LSo5Dq0t8ppZFF/deliveries/" + id
	req, err := postmates_base.New().Get(uri).Request()
	if err != nil {
		return nil, err
	}

	response := new(models.DeliveryParse)
	_, err = postmates_base.Do(req, &response, err)
	if err != nil {
		return nil, err
	}

	d := new(models.Delivery)
	d.Parse(response)

	return d, nil
}

type listResponse struct {
	Count 			int						`json:"count"`
	Payload			[]*models.Delivery		`json:"payload"`
}

func GetDeliveries() (*listResponse, error) {

	// TODO: implement metadata in the response.. already in initial response

	postmates_base := makeBaseUrl()
	uri := "/v1/customers/cus_LSo5Dq0t8ppZFF/deliveries"
	req, err := postmates_base.New().Get(uri).Request()
	if err != nil {
		return nil, err
	}

	response := new(map[string]interface{})
	_, err = postmates_base.Do(req, &response, err)
	if err != nil {
		return nil, err
	}

	var payload []interface{}
	var total_count int
	for k, v := range(*response) {
		if k == "data" {
			payload = v.([]interface{})
		} else if k == "total_count" {
			total_count = int(v.(float64))
		}
	}

	var deliveries_array []*models.Delivery
	for i := range(payload) {
		d := payload[i].(map[string]interface{})
		dp, _ := models.ParseDeliveryParse(d)

		delivery := new(models.Delivery)
		if delivery.Parse(dp); err != nil {
			return nil, err
		}

		deliveries_array = append(deliveries_array, delivery)
	}

	deliveries := new(listResponse)
	deliveries.Payload = deliveries_array
	deliveries.Count = total_count

	return deliveries, nil
}
