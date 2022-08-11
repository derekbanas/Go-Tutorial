package app2

import (
	"fmt"
	"regexp"
)

func IsEmail(s string) (string, error) {
	// Used a raw string here so I didn't have
	// to double backslashes
	r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)

	if r.MatchString(s) {
		return "Valid Email", nil
	} else {
		return "", fmt.Errorf("not a valid email")
	}
}
