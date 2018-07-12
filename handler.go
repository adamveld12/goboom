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

	if referer := req.Header.Get("Referer"); referer != "" {
		result.Referer = referer
	} else {
		result.Referer = fmt.Sprintf("http://%s", req.RemoteAddr)
	}

	if metricURL, ok := result.Metrics["u"]; ok && len(metricURL) > 0 {
		result.Source = metricURL[0]
	} else {
		result.Source = fmt.Sprintf("http://%s", req.RemoteAddr)
	}

	if ua, ok := req.Header["User-Agent"]; ok && len(ua) > 0 && ua[0] != "" {
		result.UserAgent = ua[0]
	} else {
		result.UserAgent = ""
	}

	return result, nil
}
