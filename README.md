# go_cake_api_mix
A basic API for testing with.
It's part of three API's that pass JSON around and manipulate it.

It has the following endpoints:

* PUT ingredient - adds an ingredient to a bowl
* POST bowl - creates a new bowl by name
* GET bowl - returns bowl by name

# Dependencies
golang's dep is used for dependency management:
https://github.com/golang/dep
```
dep ensure
```

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
## POST /bowl
```
curl -i -X POST localhost:8000/bowl -H 'Content-Type: application/json' -d '{"Name": "test"}'
```

## GET /bowl
```
curl -i -X GET localhost:8000/bowl/test
```

## PUT /ingredient
```
curl -i -X PUT localhost:8000/ingredient -H 'Content-Type: application/json' -d '{"BowlName": "test", "Name": "eggs", "Quantity": "all the fucking eggs"}'
```

# Docker
If you want to build this into a docker image you must ensure that all Dependencies are available and that it's compiled for linux:
```
dep ensure
env GOOS=linux GOARCH=arm go build -v
docker image build . --tag go_cake_api_mix
```

Run the image as a container:
```
docker run -d -p 8080:8000 --name mix go_cake_api_mix
```
Note: that this exposes 8000 from the container to 8080 locally, so the curl commands examples above require changing the port to 8080

Check logs:
```
docker logs mix
```

Kill and remove container:
```
docker kill mix && docker rm mix
```
