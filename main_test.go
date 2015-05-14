package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeKey(t *testing.T) {
	if len(makeKey()) != 10 {
		t.Errorf("makeKey() didn't return a string of length 10")
	}
}

func TestShortenResponse(t *testing.T) {
	request, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	recorder := httptest.NewRecorder()

	shortenResponse(recorder, request, "http://example.com/")

	fmt.Printf("%d - %s", recorder.Code, recorder.Body.String())
}

func TestIndexHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	recorder := httptest.NewRecorder()

	indexHandler(recorder, request)

	fmt.Printf("%d - %s", recorder.Code, recorder.Body.String())
}
