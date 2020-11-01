# GO backend API

API's for our go-backend:

### /go/sha256

> this API gets two numbers, add them together and returns the sum in sha256 hash algorithm
```
method: POST
request body: json {
                    first number,
                    second number,
                   }

response bondy: json {
                      result string,
                     }
```
errors:

```should any of the request body keys contain none numerical character ==> 422```

### /go/write
> this API gets a number (i.e. I) and returns the I'th line in a specific file in our servers (1<=I<=100)

```
method: GET
request body: NULL
request param: number

response bondy: raw string
```
errors:

```should the number param be greater than 100 ==> 405```

