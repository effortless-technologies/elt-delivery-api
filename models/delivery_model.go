package models

import (
	"fmt"
)

type Delivery struct {
	Kind				string				`json:"kind,omitempty"`
	Id 					string				`json:"id,omitempty"`
	Created				string 				`json:"created,omitempty"`
	Updated				string				`json:"updated,omitempty"`
	Status 				string 				`json:"status,omitempty"`
	Complete 			bool				`json:"complete,omitempty"`
	PickupEta 			string				`json:"pickup_eta,omitempty"`
	DropoffEta			string				`json:"dropoff_eta,omitempty"`
	DropoffDeadline		string				`json:"dropoff_deadline,omitempty"`
	QuoteId				string				`json:"quote_id,omitempty"`
	Fee 				int 				`json:"fee,omitempty"`
	Currency 			string				`json:"currency,omitempty"`
	Manifest 			*Manifest			`json:"manifest,omitempty"`
	DropoffIdentifier	string				`json:"dropoff_identifier,omitempty"`
	Pickup 				*Address			`json:"pickup,omitempty"`
	Dropoff				*Address			`json:"dropoff,omitempty"`
	Courier 			*Courier			`json:"courier,omitempty"`
	RelatedDeliveries	*RelatedDeliveries	`json:"related_deliveries,omitempty"`
}

type DeliveryParse struct {
	Kind				string				`json:"kind,omitempty"`
	Id 					string				`json:"id,omitempty"`
	Created				string 				`json:"created,omitempty"`
	Updated				string				`json:"updated,omitempty"`
	Status 				string 				`json:"status,omitempty"`
	Complete 			bool				`json:"complete,omitempty"`
	PickupEta 			string				`json:"pickup_eta,omitempty"`
	DropoffEta			string				`json:"dropoff_eta,omitempty"`
	DropoffDeadline		string				`json:"dropoff_deadline,omitempty"`
	QuoteId				string				`json:"quote_id,omitempty"`
	Fee 				int 				`json:"fee,omitempty"`
	Currency 			string				`json:"currency,omitempty"`
	Manifest 			interface{}			`json:"manifest,omitempty"`
	DropoffIdentifier	string				`json:"dropoff_identifier,omitempty"`
	Pickup 				interface{}			`json:"pickup,omitempty"`
	Dropoff				interface{}			`json:"dropoff,omitempty"`
	Courier 			interface{}			`json:"courier,omitempty"`
	RelatedDeliveries	interface{}			`json:"related_deliveries,omitempty"`
}

func (d *Delivery) Parse(dp *DeliveryParse) error {

	d.Kind = dp.Kind
	d.Id = dp.Id
	d.Created = dp.Created
	d.Updated = dp.Updated
	d.Status = dp.Status
	d.Complete = dp.Complete
	d.PickupEta = dp.PickupEta
	d.DropoffEta = dp.DropoffEta
	d.DropoffDeadline = dp.DropoffDeadline
	d.QuoteId = dp.QuoteId
	d.Fee = dp.Fee
	d.Currency = dp.Currency

	m := new(Manifest)
	m.Parse(dp.Manifest.(map[string]interface{}))
	d.Manifest = m

	d.DropoffIdentifier = dp.DropoffIdentifier

	pa := new(Address)
	pa.Parse(dp.Pickup.(map[string]interface{}))
	d.Pickup = pa

	da := new(Address)
	da.Parse(dp.Dropoff.(map[string]interface{}))
	d.Dropoff = da

	c := new(Courier)
	c.Parse(dp.Manifest.(map[string]interface{}))
	d.Courier = c

	return nil
}

func (d *Delivery) Print() {

	fmt.Println()
	fmt.Println("Kind: ", d.Kind)
	fmt.Println("Id: ", d.Id)
	fmt.Println("Created: ", d.Created)
	fmt.Println("Updated: ", d.Updated)
	fmt.Println("Status: ",  d.Status)
	fmt.Println("Complete: ", d.Complete)
	fmt.Println("PickupEta: ", d.PickupEta)
	fmt.Println("DropoffEta: ", d.DropoffEta)
	fmt.Println("DropoffDeadline: ", d.DropoffDeadline)
	fmt.Println("QuoteId: ", d.QuoteId)
	fmt.Println("Fee: ", d.Fee)
	fmt.Println("Currency: ", d.Currency)

	fmt.Println("Manifest: ")
	d.Manifest.Print()

	fmt.Println("DropoffIdentifier: ", d.DropoffIdentifier)

	fmt.Println("Pickup: ")
	d.Pickup.Print()

	fmt.Println("Dropoff: ")
	d.Dropoff.Print()

	fmt.Println("Courier")
	d.Courier.Print()

	fmt.Println()
}


type Manifest struct {
	Reference 			string 				`json:"reference,omitempty"`
	Description 		string 				`json:"description,omitempty"`
}

func (m *Manifest) Parse(payload map[string]interface{}) {

	for k, v := range(payload) {
		if k == "reference" {
			m.Reference = v.(string)
		} else if k == "description" {
			m.Description = v.(string)
		}
 	}
}

func (m *Manifest) Print() {

	fmt.Println("	Reference: ", m.Reference)
	fmt.Println("	Description: ", m.Description)
}

type Address struct {
	Name 				string				`json:"name,omitempty"`
	PhoneNumber			string 				`json:"phone_number,omitempty"`
	Address 			string 				`json:"address,omitempty"`
	DetailedAddress		*DetailedAddress	`json:"detailed_address,omitempty"`
	Notes				string				`json:"notes,omitempty"`
	Location 			*Location 			`json:"location,omitempty"`
}

//type AddressParse struct {
//	Name 				string				`json:"name,omitempty"`
//	PhoneNumber			string 				`json:"phone_number,omitempty"`
//	Address 			string 				`json:"address,omitempty"`
//	DetailedAddress		interface{}			`json:"detailed_address,omitempty"`
//	Notes				string				`json:"notes,omitempty"`
//	Location 			interface{}			`json:"location,omitempty"`
//}

func (a *Address) Parse(payload map[string]interface{}) {

	for k, v := range(payload) {
		if k == "name" {
			a.Name = v.(string)
		} else if k == "phone_number" {
			a.PhoneNumber = v.(string)
		} else if k == "address" {
			a.Address = v.(string)
		} else if k == "detailed_address" {
			da := new(DetailedAddress)
			da.Parse(v.(map[string]interface{}))
			a.DetailedAddress = da
		} else if k == "notes" {
			a.Notes = v.(string)
		} else if k == "location" {
			l := new(Location)
			l.Parse(v.(map[string]interface{}))
			a.Location = l
		}
	}
}

func (a *Address) Print() {

	// TODO fix printing below

	fmt.Println("	Name: ", a.Name)
	fmt.Println("	Phone Number: ", a.PhoneNumber)
	fmt.Println("	Address: ", a.Address)
	fmt.Println("	Detailed Address: ")
	a.DetailedAddress.Print()
	fmt.Println("	Notes: ", a.Notes)
	fmt.Println("	Location: ")
	a.Location.Print()
}

type DetailedAddress struct {
	StreetAddress1		string				`json:"street_address_1,omitempty"`
	StreetAddress2		string				`json:"street_address_2,omitempty"`
	City 				string				`json:"city,omitempty"`
	State 				string 				`json:"state,omitempty"`
	ZipCode				string				`json:"zip_code,omitempty"`
	Country 			string				`json:"country,omitempty"`
}

func (da *DetailedAddress) Parse(payload map[string]interface{}) {

	for k, v := range(payload) {
		if k == "street_address_1" {
			da.StreetAddress1 = v.(string)
		} else if k == "street_address_2" {
			da.StreetAddress2 = v.(string)
		} else if k == "city" {
			da.City = v.(string)
		} else if k == "state" {
			da.State = v.(string)
		} else if k == "zip_code" {
			da.ZipCode = v.(string)
		} else if k == "country" {
			da.Country = v.(string)
		}
	}
}

func (da *DetailedAddress) Print() {

	fmt.Println("		Street Address 1: ", da.StreetAddress1)
	fmt.Println("		Street Address 2: ", da.StreetAddress2)
	fmt.Println("		City: ", da.City)
	fmt.Println("		State: ", da.State)
	fmt.Println("		ZipCode: ", da.ZipCode)
	fmt.Println("		Country: ", da.Country)
}

type Courier struct {
	Name 				string 				`json:"name,omitempty"`
	Rating				string				`json:"rating,omitempty"`
	VehicleType			string				`json:"vehicle_type,omitempty"`
	PhoneNumber			string				`json:"phone_number,omitempty"`
	Location 			*Location			`json:"location,omitempty"`
	ImgHref				string 				`json:"img_href,omitempty"`
}

//type CourierParse struct {
//	Name 				string 				`json:"name,omitempty"`
//	Rating				string				`json:"rating,omitempty"`
//	VehicleType			string				`json:"vehicle_type,omitempty"`
//	PhoneNumber			string				`json:"phone_number,omitempty"`
//	Location 			interface{}			`json:"location,omitempty"`
//	ImgHref				string 				`json:"img_href,omitempty"`
//}

func (c *Courier) Parse(payload map[string]interface{}) {

	//TODO: implement Location properly

	for k, v := range(payload) {
		if k == "name" {
			c.Name = v.(string)
		} else if k == "rating" {
			c.Rating = v.(string)
		} else if k == "vehicle_type" {
			c.VehicleType = v.(string)
		} else if k == "phone_number" {
			c.PhoneNumber = v.(string)
		//} else if k == "location" {
		//	if reflect.ValueOf(v).IsNil() {
		//		l := new(Location)
		//		l.Parse(v.(map[string]interface{}))
		//		c.Location = l
		//	}
		} else if k == "img_href" {
			c.ImgHref = v.(string)
		}
	}
}

func (c *Courier) Print() {

	fmt.Println("	Name: ", c.Name)
	fmt.Println("	Rating: ", c.Rating)
	fmt.Println("	Vehicle Type: ", c.VehicleType)
	fmt.Println("	Phone Number: ", c.PhoneNumber)
	fmt.Println("	Location: ", c.Location)
	fmt.Println("	Img Href: ", c.ImgHref)
}

type Location struct {
	Lat 				float64 			`json:"lat"`
	Lng					float64				`json:"lng"`
}

func (l *Location) Parse(payload map[string]interface{}) {

	for k, v := range(payload) {
		if k == "lat" {
			l.Lat = v.(float64)
		} else if k == "lng" {
			l.Lng = v.(float64)
		}
	}
}

func (l *Location) Print() {

	fmt.Println("		Lat: ", l.Lat)
	fmt.Println("		Lng: ", l.Lng)
}

type RelatedDeliveries struct {
	Id 					string 				`json:"id,omitempty"`
	Relationship 		string 				`json:"relationship,omitempty"`
}

func (rd *RelatedDeliveries) Parse(payload map[string]interface{}) {

	// TODO: refactor: response is a list / not using these currently

	for k, v := range(payload) {
		if k == "id" {
			rd.Id = v.(string)
		} else if k == "relationship" {
			rd.Relationship = v.(string)
		}
	}
}

func (rd *RelatedDeliveries) Print() {

	fmt.Println("	Id: ", rd.Id)
	fmt.Println("	Relationship: ", rd.Relationship)
}
