all: gen build

gen:
	wire

build:
	go build
