
all: build

test:
	go test ./cmd/blil

build:
	go build ./cmd/blil
