build:
	go build -o ./bin/main .

test:
	go test -run ''

run:
	./bin/main

all: build test run