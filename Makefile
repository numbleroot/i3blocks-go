.PHONY: all clean deps build public-ip internal-ip uptime load-average temperature battery date-time volume install dir copy

all: clean build

clean:
	go clean -i ./...

deps:
	go get -t ./...

build: public-ip internal-ip uptime load-average temperature battery date-time volume

public-ip:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' public-ip.go

internal-ip:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' internal-ip.go

uptime:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' uptime.go

load-average:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' load-average.go

temperature:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' temperature.go

battery:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' battery.go

date-time:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' date-time.go

volume:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' volume.go

install: build dir copy

dir:
	mkdir -p ~/.config/i3blocks-go

copy:
	cp public-ip ~/.config/i3blocks-go/
	cp internal-ip ~/.config/i3blocks-go/
	cp uptime ~/.config/i3blocks-go/
	cp load-average ~/.config/i3blocks-go/
	cp temperature ~/.config/i3blocks-go/
	cp battery ~/.config/i3blocks-go/
	cp date-time ~/.config/i3blocks-go/
	cp volume ~/.config/i3blocks-go/
