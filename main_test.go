package main

import "testing"

func TestMakeKey(t *testing.T) {
	if len(makeKey()) != 10 {
		t.Errorf("makeKey() didn't return a string of length 10")
	}
}
