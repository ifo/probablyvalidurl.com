package main

import (
	"io"
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
	request, err, recorder := createRequestAndRecorder("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	shortenResponse(recorder, request, "http://example.com/")
}

func TestIndexHandler(t *testing.T) {
	request, err, recorder := createRequestAndRecorder("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	indexHandler(recorder, request)
}

func createRequestAndRecorder(method, url string, body io.Reader) (*http.Request, error, *httptest.ResponseRecorder) {
	request, err := http.NewRequest(method, url, body)
	recorder := httptest.NewRecorder()
	return request, err, recorder
}
