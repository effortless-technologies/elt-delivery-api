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
