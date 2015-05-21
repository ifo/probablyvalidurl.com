package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const exampleurl = "http://example.com/"

func TestShortenResponse(t *testing.T) {
	request, err, recorder := createRequestAndRecorder("GET", exampleurl, nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	shortenResponse(recorder, request, "http://example.com/")

	// exampleurl length + key length + extra info length
	if len(recorder.Body.String()) != (len(exampleurl) + 10 + 1957) {
		t.Errorf("Shortened response was not of the expected length")
	}
}

func TestShortenResponseCode(t *testing.T) {
	request, err, recorder := createRequestAndRecorder("GET", exampleurl, nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	shortenResponse(recorder, request, "http://example.com/")

	if recorder.Code != 200 {
		t.Errorf("Response code was not 200")
	}
}

func TestIndexHandler(t *testing.T) {
	request, err, recorder := createRequestAndRecorder("GET", exampleurl, nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	indexHandler(recorder, request)

	doctype := strings.Split(recorder.Body.String(), "\n")
	if doctype[0] != "<!DOCTYPE HTML>" {
		t.Errorf("Recorder body did not start with Doctype. Instead got: %q", recorder.Body.String())
	}
}

func TestIndexHandlerCode(t *testing.T) {
	request, err, recorder := createRequestAndRecorder("GET", exampleurl, nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	indexHandler(recorder, request)

	if recorder.Code != 200 {
		t.Errorf("Response code was not 200")
	}
}

func createRequestAndRecorder(method, url string, body io.Reader) (*http.Request, error, *httptest.ResponseRecorder) {
	request, err := http.NewRequest(method, url, body)
	recorder := httptest.NewRecorder()
	return request, err, recorder
}
