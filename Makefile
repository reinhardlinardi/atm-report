.PHONY: setup dev tidy gen doc
all: dev

setup:
	wire
	go mod tidy
	go build

dev:
	go build

tidy:
	go mod tidy

gen:
	wire

doc:
	swag init -d app/server -g handler.go
