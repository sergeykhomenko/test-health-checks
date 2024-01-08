# Health checks API

## Flags

**port** - port for server to listen

**timeout** - request timeout in milliseconds. Urls that responds longer than this time would be returned with `timeout` status

**debug** - set to run in debug mode

## Build

1. Try to build or run using default `go run ./cmd/server/main.go --port=8000 --timeout=5000`
2. If you face any setup-related issues, you can use Docker with the instructions in this doc.

## Build with Docker

1. Run docker ???
2. Build the image
```shell
docker build -t test-health-checks:latest .
```
3. Run the image you just created. 
```shell
docker run -it --rm -p 127.0.0.1:8000:8000 test-health-checks:latest
```
4. If you need to adjust run parameters, you can do it changing dockerfile for now

## Run

```http request
POST http://127.0.0.1:8000/ping-urls 

{
    "urls": [
        "https://google.com"
    ],
    "strategy": "first_to_fall"
}
```

Strategies that can be applied: 
 - `first_to_fall` - terminate on first inactive/timeout url
 - `at_least_one` - wait for all responses. If at least one is working, we are waiting for it