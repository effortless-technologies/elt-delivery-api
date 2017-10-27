# API spec

## GET/delivery-quote?...

This endpoint is get create a delivery quote.

ex:

```
$ curl http://localhost:1323/delivery-quote...
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
