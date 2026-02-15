.DEFAULT_GOAL := run

build:
	go build -o integrator cmd/main.go

run: build
	./integrator

up:
	go mod tidy
	go mod vendor

clean:
	rm -f integrator

test:
	go test ./...