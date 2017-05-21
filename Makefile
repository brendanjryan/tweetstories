deafult: build

vet:
	go tool vet .

build:
	go build

test:
	go test ./...

run:
	make && ./tweetstories