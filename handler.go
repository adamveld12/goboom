package goboom

import (
	"fmt"
	"net/http"
)

// Beacon represents a boomerang beacon
type Beacon struct {
	Referer   string
	Source    string
	UserAgent string
	Metrics   Metric
}

// Metric is a map of they beacon's metrics payload
type Metric map[string]string

// BeaconValidator validates a request, returning an error if the request should not be handled
type BeaconValidator func(*http.Request) error

// BeaconExporter allows for exporting the beacon or other http.Request info to various backends or services
type BeaconExporter func(*http.Request, Beacon) error

// Handler is the beacon http.Handler
type Handler struct {
	Validator BeaconValidator
	Exporter  BeaconExporter
}

func (g Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "POST" {
		http.Error(res, fmt.Sprintf("Method not allowed: %s", req.Method), http.StatusMethodNotAllowed)
		return
	}

	if g.Validator != nil {
		if err := g.Validator(req); err != nil {
			http.Error(res, fmt.Sprintf("Request invalid: %v", err), http.StatusForbidden)
			return
		}
	}

	beacon, err := parseBeacon(req)
	if err != nil {
		http.Error(res, fmt.Sprintf("could not parse beacon: %v", err), http.StatusForbidden)
		return
	}

	exporter := ConsoleExporter(nil)
	if g.Exporter != nil {
		exporter = g.Exporter
	}

	if err := exporter(req, beacon); err != nil {
		http.Error(res, fmt.Sprintf("could not export beacon: %v", err), http.StatusForbidden)
		return
	}
}

func parseBeacon(req *http.Request) (Beacon, error) {
	if err := req.ParseForm(); err != nil {
		return Beacon{}, fmt.Errorf("Request invalid: %v", err)
	}

	var result Beacon
	if len(req.Form) > 0 {
		result.Metrics = Metric{}

		for k, v := range req.Form {
			if len(v) > 0 {
				result.Metrics[k] = v[0]
			}
		}
	}

	if metricURL, ok := result.Metrics["r"]; ok {
		result.Referer = metricURL
	} else if req.Referer() != "" {
		result.Referer = req.Referer()
	}

	if sourceURL, ok := result.Metrics["u"]; ok && sourceURL != "" {
		result.Source = sourceURL
	} else if origin := req.Header.Get("Origin"); origin != "" {
		result.Source = origin
	}

	if req.UserAgent() != "" {
		result.UserAgent = req.UserAgent()
	} else {
		result.UserAgent = ""
	}

	return result, nil
}
