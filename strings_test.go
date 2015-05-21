package main

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	setupStrings()
	// random string of length 10
	length := 10
	check := randomString(length)
	if len(check) != length {
		t.Errorf("randomString(%q) wasn't of the right length", length)
	}

	// random string of length 1957
	length = 1957
	check = randomString(length)
	if len(check) != length {
		t.Errorf("randomString(%q) wasn't of the right length", length)
	}
}

func TestRandomStringInvalidChars(t *testing.T) {
	setupStrings()
	const invalidChars = "-_"
	randString := randomString(256)
	for _, v := range invalidChars {
		if strings.Contains(randString, string(v)) {
			t.Errorf("randomString(256) contains invalid character %q", v)
		}
	}
}

func TestRandomStringValidChars(t *testing.T) {
	setupStrings()
	randString := randomString(256)
	for _, v := range randString {
		if !strings.Contains(string(alphabet), string(v)) {
			t.Errorf("randomString(256) contains invalid character %q", v)
		}
	}
}
