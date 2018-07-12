all: clean goboom dev

clean:
	rm -rf ./goboom

ci: clean
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o ./goboom ./cli

goboom:
	go build -o ./goboom ./cli

dev: clean goboom
	./goboom \
		-address localhost:8000 \
	  	-origin .*\.liveauctioneers\.com \
		-path /beacon

.PHONY: all ci clean dev