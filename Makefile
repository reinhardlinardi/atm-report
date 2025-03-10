all: dev

setup:
	wire
	go mod tidy
	go build

dev:
	wire
	go build
	go mod tidy
