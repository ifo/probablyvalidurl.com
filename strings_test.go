package main

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	// random string of length 10
	length := 10
	check := randomString(length)
	if len(check) != length {
		t.Errorf("RandomString(%q) wasn't of the right length", length)
	}

	// random string of length 1957
	length = 1957
	check = randomString(length)
	if len(check) != length {
		t.Errorf("RandomString(%q) wasn't of the right length", length)
	}
}

func TestRandomStringInvalidChars(t *testing.T) {
	const invalidChars = "-_"
	randString := randomString(256)
	for _, v := range invalidChars {
		if strings.Contains(randString, string(v)) {
			t.Errorf("RandomString(256) contains invalid character %q", v)
		}
	}
}

func TestRandomStringValidChars(t *testing.T) {
	randString := randomString(256)
	for _, v := range randString {
		if !strings.Contains(alphabet, string(v)) {
			t.Errorf("RandomString(256) contains invalid character %q", v)
		}
	}
}
