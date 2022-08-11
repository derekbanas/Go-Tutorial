package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("derek@aol.com")
	if err != nil {
		t.Error("derek@aol.com is an email")
	}

	_, err = IsEmail("derek@aol")
	if err != nil {
		t.Error("derek@aol is not email")
	}
}
