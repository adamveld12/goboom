all: clean goboom dev

clean:
	rm -rf ./goboom

ci: clean
	go get -t -d -v ./... 
	go test -v -cover ./...
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o ./goboom ./cli

goboom:
	go build -o ./goboom ./cli

dev: clean goboom
	./goboom \
		-address localhost:3000 \
	  	-origin .* \
		-url "/beacon"

.PHONY: all ci clean dev