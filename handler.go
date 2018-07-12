package goboom

import (
	"fmt"
	"net/http"
)

type Beacon struct {
	Referer   string
	Source    string
	UserAgent string
	Metrics   Metric
}

type Metric map[string][]string

type BeaconValidator func(*http.Request) error
type BeaconExporter func(*http.Request, Beacon) error

type Goboom struct {
	Method    string
	URL       string
	Validator BeaconValidator
	Exporter  BeaconExporter
}

func (g Goboom) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if g.Method != "" && req.Method != g.Method {
		http.Error(res, "Unexpected http.method", http.StatusMethodNotAllowed)
		return
	}

	if g.URL != "" && req.URL.Path != g.URL {
		http.Error(res, "Not Found", http.StatusNotFound)
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
			result.Metrics[k] = v
		}
	}

	if metricURL, ok := result.Metrics["r"]; ok && len(metricURL) > 0 {
		result.Referer = metricURL[0]
	} else if req.Referer() != "" {
		result.Referer = req.Referer()
	}

	if sourceURL, ok := result.Metrics["u"]; ok && len(sourceURL) > 0 && sourceURL[0] != "" {
		result.Source = sourceURL[0]
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
