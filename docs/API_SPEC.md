# API spec

## GET/deliveries/quote?...

This endpoint is get create a delivery quote.

ex:

```
$ curl http://localhost:1323/deliveries/quote...
```

#### Query Parameters

Mandatory Query Parameters

```
pickup-address={string}
dropoff-address={string}

```

#### Response
```
{
    "kind": {string},
    "id": {string},
    "created": {string},
    "expires": {string},
    "fee": {int},
    "currency": {string},
    "dropoff_eta": {string},
    "duration": {int}
}
```

## POST/deliveries/:quote_id

This endpoint is create a delivery using the quote ID created in the request
directly above.

#### Request
```
{
    "manifest": {string},                                     
    "manifest_reference": {string},                         // optional
    "pickup_name": {string},
    "pickup_address": {string},
    "pickup_phone_number": {string},
    "pickup_business_name": {string},                       // optional
    "pickup_notes": {string},                               // optional
    "dropoff_name": {string},   =
    "dropoff_address": {string},
    "dropoff_phone_number": {string},
    "dropoff_business_name": {string},                      // optional
    "dropoff_notes": {string},                              // optional
    "require_id": {string}                                  // optional
}
```

#### Response
```
{
    "kind": {string},
    "id": {string},
    "created": {string},
    "updated": {string},
    "status": {string},
    "expires": {string},
    "complete": {bool},
    "pickup_eta": {string},
    "dropoff_eta": {string},
    "dropoff_deadline": {string},
    "quote_id": {string},
    "fee": {int},
    "currency": {string},
    "manifest": {
        "reference": {string},
        "description" : {string},
    },
    "dropoff_identifier": {string},
    "pickup": {
        "name": {string},
        "phone_number": {string},
        "address:: {string},
        "detailed_address": {
            "street_address_1": {string},
            "street_address_2": {string},
            "city": {string},
            "state": {string},
            "zip_code": {string},
            "country": {string},
        },
        "notes": {string},
        "location": {
            "lat": {float64},
            "lng": {float64}
        }
    },
    "dropoff": {
        "name": {string},
        "phone_number": {string},
        "address:: {string},
        "detailed_address": {
            "street_address_1": {string},
            "street_address_2": {string},
            "city": {string},
            "state": {string},
            "zip_code": {string},
            "country": {string},
        },
        "notes": {string},
        "location": {
            "lat": {float64},
            "lng": {float64}
        }
    },
    "courier": {
        "name": {string},
        "rating": {string},
        "vehicle_type": {string},
        "phone_number": {string},
        "location": {
            "lat": {float64},
            "lng": {float64}
        }
        "img_href": {string}
    }
}
```

## GET/deliveries/:delivery_id

This endpoint is get a specific delivery with an delivery id.

ex:

```
$ curl http://localhost:1323/deliveries/:delivery_id
```

#### Response
```
{
    "kind": {string},
    "id": {string},
    "created": {string},
    "updated": {string},
    "status": {string},
    "expires": {string},
    "complete": {bool},
    "pickup_eta": {string},
    "dropoff_eta": {string},
    "dropoff_deadline": {string},
    "quote_id": {string},
    "fee": {int},
    "currency": {string},
    "manifest": {
        "reference": {string},
        "description" : {string},
    },
    "dropoff_identifier": {string},
    "pickup": {
        "name": {string},
        "phone_number": {string},
        "address:: {string},
        "detailed_address": {
            "street_address_1": {string},
            "street_address_2": {string},
            "city": {string},
            "state": {string},
            "zip_code": {string},
            "country": {string},
        },
        "notes": {string},
        "location": {
            "lat": {float64},
            "lng": {float64}
        }
    },
    "dropoff": {
        "name": {string},
        "phone_number": {string},
        "address:: {string},
        "detailed_address": {
            "street_address_1": {string},
            "street_address_2": {string},
            "city": {string},
            "state": {string},
            "zip_code": {string},
            "country": {string},
        },
        "notes": {string},
        "location": {
            "lat": {float64},
            "lng": {float64}
        }
    },
    "courier": {
        "name": {string},
        "rating": {string},
        "vehicle_type": {string},
        "phone_number": {string},
        "location": {
            "lat": {float64},
            "lng": {float64}
        }
        "img_href": {string}
    }
}
```
