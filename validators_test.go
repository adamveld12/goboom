package goboom

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"
)

type originValidatorCase struct {
	Name           string
	InputRegex     string
	InputOrigin    string
	ExpectedError  error
	ExpectedOutput error
}

func TestOriginValidator(t *testing.T) {
	cases := []originValidatorCase{
		{
			Name:           "Happy Path",
			InputRegex:     ".*",
			InputOrigin:    "www.example.com",
			ExpectedError:  nil,
			ExpectedOutput: nil,
		},
		{
			Name:           "Validation passes",
			InputRegex:     ".*\\.example\\.com",
			InputOrigin:    "www.example.com",
			ExpectedError:  nil,
			ExpectedOutput: nil,
		},
		{
			Name:           "validation fails",
			InputRegex:     ".*\\.example\\.com",
			InputOrigin:    "www.google.com",
			ExpectedError:  nil,
			ExpectedOutput: fmt.Errorf("origin does not match regex '.*\\.example\\.com': got 'www.google.com'"),
		},
		{
			Name:           "validation fails empty string",
			InputRegex:     ".*\\.example\\.com",
			InputOrigin:    "",
			ExpectedError:  nil,
			ExpectedOutput: fmt.Errorf("origin does not match regex '.*\\.example\\.com': got ''"),
		},
		{
			Name:           "Validator creation fails - empty string",
			InputRegex:     "",
			InputOrigin:    "",
			ExpectedError:  errors.New("empty string not allowed"),
			ExpectedOutput: nil,
		},
		{
			Name:           "Validator creation fails - empty string",
			InputRegex:     "*",
			InputOrigin:    "",
			ExpectedError:  errors.New("empty string not allowed"),
			ExpectedOutput: nil,
		},
	}

	for idx, c := range cases {
		v, err := OriginValidator(c.InputRegex)
		if !reflect.DeepEqual(err, c.ExpectedError) {
			t.Errorf("#%d %s error creating validator\nexpected:\n%v\ngot:\n%v", idx, c.Name, c.ExpectedError, err)
			return
		} else if c.ExpectedError != nil {
			return
		}

		inputURL := fmt.Sprintf("http://%s/beacon", c.InputOrigin)
		inputReq := httptest.NewRequest("POST", inputURL, nil)
		inputReq.Header.Set("Origin", c.InputOrigin)
		validationErr := v(inputReq)
		if !reflect.DeepEqual(validationErr, c.ExpectedOutput) {
			t.Errorf("#%d %s validation failed\nexpected:\n%v\ngot:\n%v", idx, c.Name, c.ExpectedOutput, err)
		}
	}
}
