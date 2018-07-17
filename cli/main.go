package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/adamveld12/goboom"
)

var (
	address = flag.String("address", "0.0.0.0:3000", "the http ip/port to listen on")
	url     = flag.String("url", "/beacon", "the http url that the beacon calls will come in on")
	origin  = flag.String("origin", "", "an origin regex to validate beacon requests against. If set an origin header must be present")
)

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	gb := goboom.Handler{
		Exporter: goboom.ConsoleExporter(os.Stdout),
	}

	if *origin != "" {
		var err error
		gb.Validator, err = goboom.OriginValidator(*origin)
		if err != nil {
			fmt.Printf("Could not compile origin regex: %v\n", err)
			os.Exit(-1)
			return
		}
	}

	mux.Handle(*url, gb)

	fmt.Printf("Listening @ %s for HTTP calls on \"%s\"\n", *address, *url)
	if err := http.ListenAndServe(*address, mux); err != nil {
		fmt.Printf("Could not listen and serve on %s: %v", *address, err)
		os.Exit(1)
	}
}
