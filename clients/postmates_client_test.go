package clients

import (
	"testing"

	"github.com/effortless-technologies/elt-delivery-api/models"

	. "github.com/smartystreets/goconvey/convey"
)

func createTestQuote() *models.Quote {

	q := new(models.Quote)
	q.Kind = "delivery_quote"
	q.Fee = 905
	q.Currency = "usd"
	q.Duration = 40

	return q
}

func createTestDelivery() *models.Delivery {

	d := new(models.Delivery)
	d.Kind = "delivery"
	d.Status = "pending"
	d.Complete = false
	d.Currency = "usd"

	return d
}

func createParams(id string) *CreateDeliveryParams {

	return &CreateDeliveryParams{
		QuoteId: id,
		Manifest: "A box of gray kittens",
		ManifestReference: "",
		PickupName: "Laser Crazer",
		PickupAddress: "20 McAllister St, San Francisco, CA",
		PickupPhoneNumber: "415-555-4242",
		PickupBusinessName: "",
		PickupNotes: "",
		DropoffName: "Matt's Palace",
		DropoffAddress: "101 Market St, San Francisco, CA",
		DropoffPhoneNumber: "559-555-4242",
		DropoffBusinessName: "",
		DropoffNotes: "",
		RequireId: "",
	}
}

func TestPostmatesClient_GetDeliveryQuote(t *testing.T) {

	Convey("Given a test Quote", t, func() {

		test_quote := createTestQuote()
		So(test_quote, ShouldNotBeNil)

		Convey("When a test quote is requested", func() {

			quote, err := GetDeliveryQuote(
				"20 McAllister St, San Francisco, CA",
				"101 Market St, San Francisco, CA",
			)
			So(quote, ShouldNotBeNil)
			So(err,ShouldBeNil)
			Convey("Then the received quote should resemble the " +
				"test quote", func() {

				So(quote.Kind, ShouldEqual, test_quote.Kind)
				//So(quote.Fee, ShouldEqual, test_quote.Fee)
				So(quote.Currency, ShouldEqual, test_quote.Currency)
				//So(quote.Duration, ShouldEqual, test_quote.Duration)
			})
		})
	})
}

func TestPostmatesClient_CreateDelivery(t *testing.T) {

	Convey("Given a test Quote and a test Delivery", t, func() {

		test_delivery := createTestDelivery()
		So(test_delivery, ShouldNotBeNil)

		quote, err := GetDeliveryQuote(
			"20 McAllister St, San Francisco, CA",
			"101 Market St, San Francisco, CA",
		)
		So(quote, ShouldNotBeNil)
		So(err,ShouldBeNil)

		Convey("When a new Delivery is created", func() {

			params := createParams(quote.Id)
			So(params, ShouldNotBeNil)

			delivery, err := CreateDelivery(params)
			So(err, ShouldBeNil)
			So(delivery, ShouldNotBeNil)

			Convey("Then the Created Delivery should resemble the " +
				"test Delivery", func() {

				So(delivery.Kind, ShouldEqual, test_delivery.Kind)
				So(delivery.Status, ShouldEqual, test_delivery.Status)
				So(delivery.Complete, ShouldEqual, test_delivery.Complete)
				So(delivery.Currency, ShouldEqual, test_delivery.Currency)
			})
		})
	})
}

func TestPostmatesClient_GetDelivery(t *testing.T) {

	Convey("Given a test Quote and a test Delivery", t, func() {

		test_delivery := createTestDelivery()
		So(test_delivery, ShouldNotBeNil)

		quote, err := GetDeliveryQuote(
			"20 McAllister St, San Francisco, CA",
			"101 Market St, San Francisco, CA",
		)
		So(quote, ShouldNotBeNil)
		So(err,ShouldBeNil)

		Convey("When a new Delivery is created and retrieved", func() {

			params := createParams(quote.Id)
			So(params, ShouldNotBeNil)

			created_delivery, err := CreateDelivery(params)
			So(err, ShouldBeNil)
			So(created_delivery, ShouldNotBeNil)
			So(created_delivery.Kind, ShouldEqual, test_delivery.Kind)
			So(created_delivery.Status, ShouldEqual, test_delivery.Status)
			So(created_delivery.Complete, ShouldEqual, test_delivery.Complete)
			So(created_delivery.Currency, ShouldEqual, test_delivery.Currency)
			So(created_delivery.Id, ShouldNotBeNil)

			retrieved_delivery, err := GetDelivery(created_delivery.Id)

			Convey("Then the Retieved Delivery should resemble the " +
				"Test Delivery and Created Delivery", func() {

				So(retrieved_delivery.Kind, ShouldEqual, test_delivery.Kind)
				So(retrieved_delivery.Complete, ShouldEqual,
					test_delivery.Complete)
				So(retrieved_delivery.Currency, ShouldEqual,
					test_delivery.Currency)

				So(retrieved_delivery.Kind, ShouldEqual, created_delivery.Kind)
				So(retrieved_delivery.Complete, ShouldEqual,
					created_delivery.Complete)
				So(retrieved_delivery.Currency, ShouldEqual,
					created_delivery.Currency)
			})
		})
	})
}

func TestPostmatesClient_GetDeliveries(t *testing.T) {

	Convey("Given Test Deliveries", t, func() {

		Convey("When retrieving those Deliveries", func() {

			deliveries, err := GetDeliveries()
			So(deliveries, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Then the returned Deliveries should be in the " +
				"correct format", func() {

				So(deliveries.Count, ShouldNotBeNil)
				So(deliveries.Payload, ShouldNotBeNil)
			})
		})
	})
}
