package goboom

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// ConsoleExporter will write the beacons to the specified output stream, using os.Stdout if nil is passed
func ConsoleExporter(output io.Writer) BeaconExporter {
	if output == nil {
		output = os.Stdout
	}

	return func(req *http.Request, b Beacon) error {
		fmt.Fprintf(output, "%s - %s\n", req.Method, req.URL.String())
		fmt.Fprintf(output, "Source=%s\nReferer=%s\nUser-Agent=%s\nMetrics:\n", b.Source, b.Referer, b.UserAgent)

		count := 0
		for k, v := range b.Metrics {
			if len(v) > 0 {
				count++
				fmt.Fprintf(output, "\t%v=%v\n", k, v)
			}
		}

		fmt.Fprintf(output, "%v Metrics imported\n\n", count)
		return nil
	}
}
