.PHONY: all clean deps build public-ip internal-ip load-average install dir copy

all: clean build

clean:
	go clean -i ./...

deps:
	go get -t ./...

build: public-ip internal-ip load-average

public-ip:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' public-ip.go

internal-ip:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' internal-ip.go

load-average:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' load-average.go

install: build dir copy

dir:
	mkdir -p ~/.config/i3blocks-go

copy:
	cp public-ip ~/.config/i3blocks-go/
	cp internal-ip ~/.config/i3blocks-go/
	cp load-average ~/.config/i3blocks-go/