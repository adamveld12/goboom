package goboom

import (
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

type methodTestCase struct {
	Name           string
	ExpectedMethod string
	Method         string
	ExpectedStatus int
}

func TestMethodValidation(t *testing.T) {
	cases := []methodTestCase{
		{Name: "Happy Path", ExpectedMethod: "POST", Method: "POST", ExpectedStatus: 200},
		{Name: "Forbidden when not matched", ExpectedMethod: "POST", Method: "GET", ExpectedStatus: 405},
		{Name: "Accept anything when unintialized 1", ExpectedMethod: "", Method: "GET", ExpectedStatus: 200},
		{Name: "Accept anything when unintialized 2", ExpectedMethod: "", Method: "POST", ExpectedStatus: 200},
	}

	str, _ := os.Open(os.DevNull)
	defer str.Close()

	for idx, c := range cases {
		g := Goboom{
			Method:   c.ExpectedMethod,
			URL:      "/",
			Exporter: ConsoleExporter(str),
		}

		req := httptest.NewRequest(c.Method, "http://127.0.0.1:3000/", nil)
		w := httptest.NewRecorder()
		g.ServeHTTP(w, req)
		res := w.Result()
		if res.StatusCode != c.ExpectedStatus {
			t.Errorf("Method test case #%d failed: expected '%d' got '%d'", idx, c.ExpectedStatus, res.StatusCode)
		}
	}
}

type urlTestCase struct {
	Name           string
	InputURL       string
	TestURL        string
	ExpectedStatus int
}

func TestURLValidation(t *testing.T) {
	cases := []urlTestCase{
		{Name: "Happy Path", InputURL: "/beacon", TestURL: "/beacon", ExpectedStatus: 200},
		{Name: "Path doesn't match", InputURL: "/not-beacon-url", TestURL: "/beacon", ExpectedStatus: 404},
		{Name: "Not initialized allows anything #1", InputURL: "", TestURL: "/beacon", ExpectedStatus: 200},
		{Name: "Not initialized allows anything #2", InputURL: "", TestURL: "/beacon-url-2", ExpectedStatus: 200},
	}
	str, _ := os.Open(os.DevNull)
	defer str.Close()

	for idx, c := range cases {
		g := Goboom{
			URL:      c.InputURL,
			Exporter: ConsoleExporter(str),
		}

		req := httptest.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:3000%s", c.TestURL), nil)
		w := httptest.NewRecorder()
		g.ServeHTTP(w, req)
		res := w.Result()
		if res.StatusCode != c.ExpectedStatus {
			t.Errorf("URL test case #%d - %s failed: expected '%d' got '%d'",
				idx,
				c.Name,
				c.ExpectedStatus,
				res.StatusCode)
		}
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
					"r":        []string{"http://boomerang-test.surge.sh/"},
					"u":        []string{"http://boomerang-test.surge.sh/test"},
					"c.tti.vr": []string{"665"},
				},
			},
		},
		//{Name: "Parses user agent correctly"},
	}

	for idx, c := range cases {
		req := httptest.NewRequest("POST", "http://127.0.0.1/beacon", c.InputBody)
		req.Header.Add("Origin", c.InputSourcePage)
		req.Header.Add("Referer", c.InputReferer)
		req.Header.Add("User-Agent", c.InputUserAgent)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.RemoteAddr = "127.0.0.1:80"
		b, err := parseBeacon(req)

		if err != c.ExpectedErr {
			t.Errorf("parseBeacon case #%d - %s failed: expected '%v' got '%v'",
				idx,
				c.Name,
				c.ExpectedErr,
				err)
			return
		}

		if !reflect.DeepEqual(b, c.ExpectedBeacon) {
			t.Errorf("parseBeacon case #%d - %s failed: beacon \nexpected\n'%+v' \ngot\n'%+v'",
				idx,
				c.Name,
				c.ExpectedBeacon,
				b)
			return
		}
	}
}

var postBuf = bytes.NewBufferString("u=http%3A%2F%2Fboomerang-test.surge.sh%2Ftest&r=http%3A%2F%2Fboomerang-test.surge.sh%2F&c.tti.vr=665")
