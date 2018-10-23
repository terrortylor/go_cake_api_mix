# go_cake_api_mix
A basic API for testing with.
It's part of three API's that pass JSON around and manipulate it.

It has the following endpoints:

* PUT ingredient - adds an ingredient to a bowl
* DELETE ingredient - deletes an ingredient from a bowl
* POST bowl - creates a new bowl by name
* GET bowl - returns bowl by name

# Dependencies
golang's dep is used for dependency management:
https://github.com/golang/dep
'''
dep ensure
'''

# Run
Either just run it:
```
go run main.go
```

or compile and then run executable:
```
go build
./go_cake_api_mix
```

# cURL examples:
## POST bowl
```
curl -i -X POST localhost:8000/bowl -H 'Content-Type: application/json' -d '{"Name": "test"}'
```

## GET bowl
```
curl -i -X GET localhost:8000/bowl/test
```
