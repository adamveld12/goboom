package goboom

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

// OriginValidator takes a regex and returns a validator that checks the origin header of the request against the regex
func OriginValidator(originRegex string) (BeaconValidator, error) {
	if originRegex == "" {
		return nil, errors.New("empty string not allowed")
	}

	validatorRegex, err := regexp.Compile(originRegex)
	if err != nil {
		return nil, fmt.Errorf("could not compile regex string '%s': %v", originRegex, err)
	}

	return func(req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" || !validatorRegex.MatchString(origin) {
			return fmt.Errorf("origin does not match regex '%s': got '%v'", originRegex, origin)
		}
		return nil
	}, nil
}
