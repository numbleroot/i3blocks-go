.PHONY: all clean deps build

all: clean build

clean:
	go clean -i ./...

deps:
	go get -t ./...

build:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"'