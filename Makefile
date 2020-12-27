deafult: build

vet:
	go vet .

build:
	go build

test:
	go test ./...

run:
	make && ./tweetstories
