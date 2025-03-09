all: build

build:
	wire
	go build
	go mod tidy
