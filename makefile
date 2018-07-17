.PHONY: ci clean dev lint setup test

dev: clean goboom
	./goboom \
		-address localhost:3000 \
	  	-origin .* \
		-url "/beacon"

ci: clean setup lint test goboom

pc: clean lint test goboom

clean:
	rm -rf ./goboom

setup:
	go get -t -d -v ./... 

lint:
	go get golang.org/x/lint/golint
	golint -set_exit_status
	go vet -all -v

test:
	go test -v -cover ./...


goboom:
	go build -o ./goboom ./cli