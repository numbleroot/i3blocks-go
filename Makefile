.PHONY: all clean deps build public-ip install dir copy

all: clean build

clean:
	go clean -i ./...

deps:
	go get -t ./...

build: public-ip

public-ip:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' public-ip.go

install: build dir copy

dir:
	mkdir -p ~/.config/i3blocks-go

copy:
	cp public-ip ~/.config/i3blocks-go/