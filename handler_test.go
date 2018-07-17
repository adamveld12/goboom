package goboom

import (
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

type methodTestCase struct {
	Name           string
	ExpectedMethod string
	Method         string
	ExpectedStatus int
}

func TestMethodValidation(t *testing.T) {
	cases := []methodTestCase{
		{Name: "Accept Post", Method: "POST", ExpectedStatus: 200},
		{Name: "Accept Get", Method: "GET", ExpectedStatus: 200},
		{Name: "Deny Delete", Method: "DELETE", ExpectedStatus: 405},
		{Name: "Deny Head", Method: "HEAD", ExpectedStatus: 405},
		{Name: "Deny Options", Method: "OPTIONS", ExpectedStatus: 405},
	}

	str, _ := os.Open(os.DevNull)
	defer str.Close()

	for _, c := range cases {
		t.Run(c.Name, func(tsub *testing.T) {
			g := Handler{
				Exporter: ConsoleExporter(str),
			}

			req := httptest.NewRequest(c.Method, "http://127.0.0.1:3000/", nil)
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)
			res := w.Result()
			if res.StatusCode != c.ExpectedStatus {
				tsub.Errorf("expected '%d' got '%d'", c.ExpectedStatus, res.StatusCode)
			}
		})
	}
}

type urlTestCase struct {
	Name           string
	TestURL        string
	ExpectedStatus int
}

func TestURLValidation(t *testing.T) {
	cases := []urlTestCase{
		{Name: "Happy Path", TestURL: "/beacon", ExpectedStatus: 200},
		{Name: "works with empty path", TestURL: "", ExpectedStatus: 200},
		{Name: "works with multi paths", TestURL: "/beacon/url/2", ExpectedStatus: 200},
	}
	str, _ := os.Open(os.DevNull)
	defer str.Close()

	for _, c := range cases {
		t.Run(c.Name, func(tsub *testing.T) {
			g := Handler{
				Exporter: ConsoleExporter(str),
			}

			req := httptest.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:3000%s", c.TestURL), nil)
			w := httptest.NewRecorder()
			g.ServeHTTP(w, req)
			res := w.Result()
			if res.StatusCode != c.ExpectedStatus {
				tsub.Errorf(" expected '%d' got '%d'",
					c.ExpectedStatus,
					res.StatusCode)
			}
		})
	}
}

type parseBeaconTestCase struct {
	Name            string
	InputSourcePage string
	InputReferer    string
	InputUserAgent  string
	InputBody       io.Reader
	ExpectedErr     error
	ExpectedBeacon  Beacon
}

func TestParseBeacon(t *testing.T) {
	cases := []parseBeaconTestCase{
		{
			Name:            "Happy Path",
			InputSourcePage: "http://boomerang-test.surge.sh/",
			ExpectedErr:     nil,
			ExpectedBeacon: Beacon{
				Source: "http://boomerang-test.surge.sh/",
			},
		},
		{
			Name:            "Parses beacon correctly",
			InputReferer:    "http://boomerang-test.surge.sh/",
			InputSourcePage: "http://boomerang-test.surge.sh/test",
			InputUserAgent:  "",
			InputBody:       postBuf,
			ExpectedErr:     nil,
			ExpectedBeacon: Beacon{
				Referer: "http://boomerang-test.surge.sh/",
				Source:  "http://boomerang-test.surge.sh/test",
				Metrics: Metric{
					"r":        "http://boomerang-test.surge.sh/",
					"u":        "http://boomerang-test.surge.sh/test",
					"c.tti.vr": "665",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(tsub *testing.T) {
			req := httptest.NewRequest("POST", "http://127.0.0.1/beacon", c.InputBody)
			req.Header.Add("Origin", c.InputSourcePage)
			req.Header.Add("Referer", c.InputReferer)
			req.Header.Add("User-Agent", c.InputUserAgent)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.RemoteAddr = "127.0.0.1:80"
			if c.ExpectedBeacon.RemoteIP == "" {
				c.ExpectedBeacon.RemoteIP = "127.0.0.1"
			}

			b, err := parseBeacon(req)
			if err != c.ExpectedErr {
				t.Errorf("expected '%v' got '%v'",
					c.ExpectedErr,
					err)
				return
			}

			if time.Since(b.Created) > time.Second {
				t.Errorf("beacon created time is not set")
				return
			}
			c.ExpectedBeacon.Created = b.Created

			if !reflect.DeepEqual(b, c.ExpectedBeacon) {
				t.Errorf("beacon \nexpected\n'%+v' \ngot\n'%+v'",
					c.ExpectedBeacon,
					b)
				return
			}
		})
	}
}

var postBuf = bytes.NewBufferString("u=http%3A%2F%2Fboomerang-test.surge.sh%2Ftest&r=http%3A%2F%2Fboomerang-test.surge.sh%2F&c.tti.vr=665")

type parseForwardedCase struct {
	Name     string
	Input    string
	Expected string
}

func TestParseForwarded(t *testing.T) {
	cases := []parseForwardedCase{
		{
			Name:     "Empty string",
			Input:    "",
			Expected: "",
		},
		{
			Name:     "One for entry (ipv4)",
			Input:    "for=123.34.567.89",
			Expected: "123.34.567.89",
		},
		{
			Name:     "One for entry (ipv4 w\\ port)",
			Input:    "for=123.34.567.89:1234",
			Expected: "123.34.567.89",
		},
		{
			Name:     "One for entry (ipv6)",
			Input:    "for=\"[2001:db8:cafe::17]\"",
			Expected: "2001:db8:cafe::17",
		},
		{
			Name:     "One for entry (ipv6 w\\ port)",
			Input:    "for=\"[2001:db8:cafe::17]:4711\"",
			Expected: "2001:db8:cafe::17",
		},
		{
			Name:     "2 for entry",
			Input:    "for=123.34.567.89; for=98.87.654.321",
			Expected: "123.34.567.89",
		},
		{
			Name:     "2 for entry (one ipv6)",
			Input:    "for=123.34.567.89; for=\"[2001:db8:cafe::17]\"",
			Expected: "123.34.567.89",
		},
		{
			Name:     "crazy (one ipv4)",
			Input:    "by=127.0.0.1; for=123.34.567.89; host=local.example.com; proto=http",
			Expected: "123.34.567.89",
		},
		{
			Name:     "crazy (one ipv6)",
			Input:    "by=127.0.0.1; for=\"[2001:db8:cafe::17]\"; host=local.example.com; proto=http",
			Expected: "2001:db8:cafe::17",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(tsub *testing.T) {
			actual := parseForwarded(c.Input)
			if actual != c.Expected {
				t.Errorf(" failed:\nexpected:\n%s\ngot:\n%s",
					c.Expected,
					actual)
			}
		})
	}
}
