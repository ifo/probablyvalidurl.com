package main

import "testing"

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
