## Run nodejs file

You can simply go to nodejs directory and the run the command below to run main.js .
>node main.js

# Nodejs backend API

API's for our go-backend:

### /nodejs/sha256

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

```should any of the request body keys contain none numerical character ==> 401```

### /nodejs/write
> this API gets a number (i.e. I) and returns the I'th line in a specific file in our servers (1<=I<=100)

```
method: GET
request body: number

response bondy: raw string
```
errors:

```should the number param be greater than 100 or NaN ==> 401```

