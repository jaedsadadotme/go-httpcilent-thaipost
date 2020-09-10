# HttpClient Api ThaiPost

```sh
$ go run .
```

---
### Get Token
- Method -> POST
- Authorization -> Token <token>
```sh
$ curl -X POST 'http://127.0.0.1:1323/thaiPost/getToken' \
-H 'Authorization: Token <Token>' \
-H 'Content-Type: application/json'
```

Response
```sh
{
    "Status": "200 ",
    "StatusCode": 200,
    "Datas": {
        "expire": "2020-10-10 22:51:15+07:00",
        "token": <token>
    }
}
```

> Token [Api](https://track.thailandpost.co.th/dashboard)

### Get Tracking
- Method -> POST
- Authorization -> Token < Token from /getToken >
- body 
    - status -> 103 || all
    - language -> TH , EN , CN
    - barcode -> type [tracking no] if mutiple code can use comma ``ex.[xxx, yyy]``

```sh
curl -X POST 'http://127.0.0.1:1323/thaiPost/getTracking' \
-H 'Authorization: Token < Token from /getToken >' \
-H 'Content-Type: application/json' \
-d '{ "status": "103", "language": "TH", "barcode": [ "EG972610655TH" ]}'
```

Response
```
{
    "Status": "200 ",
    "StatusCode": 200,
    "Datas": {
        "message": "successful",
        "response": {
            "items": {
                "EG972610655TH": [
                    {
                        "barcode": "EG972610655TH",
                        "delivery_datetime": null,
                        "delivery_description": null,
                        "delivery_status": null,
                        "location": "คต.เทศบาลแพรกษา",
                        "postcode": "00253",
                        "receiver_name": null,
                        "signature": null,
                        "status": "103",
                        "status_date": "10/09/2563 16:03:14+07:00",
                        "status_description": "รับฝาก"
                    }
                ]
            },
            "track_count": {
                "count_number": 37,
                "track_count_limit": 1000,
                "track_date": "10/09/2563"
            }
        },
        "status": true
    }
}
```

---
Reference
- [api thaipost](https://track.thailandpost.co.th/developerGuide)