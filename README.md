# dmv-service-gin
Facade to _dmv.ca_ appointments written in plain Go using only _stdlib_

## Running locally

Before running ensure you have Go version > 1.22.0 
```shell
go version
```

Run the service locally
```shell
go run .
```

## Usage

Provide target latitude | longitude as input parameters, f.e.

```shell
curl -X GET http://localhost:8080/test?lat=33.9911214&lon=-118.4279929
```

## TODOs
- add UI for web version
- play around with goroutines for calculations / sorting
- rework to Gin (since the name of this repo)
- add CLI version
