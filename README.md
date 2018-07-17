# Goboom

[![GoDoc](https://godoc.org/github.com/adamveld12/goboom?status.svg)](http://godoc.org/github.com/adamveld12/goboom)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/adamveld12/goboom)
[![Gocover](http://gocover.io/_badge/github.com/adamveld12/goboom)](http://gocover.io/github.com/adamveld12/goboom)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamveld12/goboom)](https://goreportcard.com/report/github.com/adamveld12/goboom)
[![Build Status](https://semaphoreci.com/api/v1/adamveld12/goboom/branches/master/badge.svg)](https://semaphoreci.com/adamveld12/goboom)


Boomerang beacon HTTP server.

You can write custom validators and exporters, allowing you to pipe the beacons into any kind of backend you wish.

## How to console

```sh
goboom -address "127.0.0.1:3000" -origin '.*\\.example\\.com' -url "/beacon"
```

## How to library

```golang
func main() {
	gb := goboom.Goboom {
		Exporter: goboom.ConsoleExporter(os.Stdout),
	}

	log.Fatal(http.ListenAndServe("127.0.0.1:3000", gb))
}
```

## How to contribute

Start the server using `make dev`


You will need a page with boomerang setup to emit beacons at your server. The simplest way is to use this page:


[http://boomerang-test.surge.sh/](http://boomerang-test.surge.sh/)


This will send an HTTP `POST` to `localhost:3000/beacon` making it easy for you to test things.

## LICENSE 

MIT