package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// TODO replace database
var sites map[string]string = make(map[string]string)

// TODO make these into env vars
const port string = ":3000"
const sitePrefix string = "http://localhost" + port

func main() {
	// setup random strings
	setup()

	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	bdy, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(bdy)
	if body != "" {
		body, _ = url.QueryUnescape(body)
		body = body[4:]
	}

	switch r.URL.Path {
	case "/":
		if body == "" {
			http.ServeFile(w, r, "index.html")
		} else {
			shortenResponse(w, body)
		}
	default:
		// we don't actually care beyond 10 characters
		if len(r.URL.Path) > 10 {
			http.Redirect(w, r, sites[r.URL.Path[1:11]], http.StatusTemporaryRedirect)
		} else {
			http.NotFound(w, r)
		}
	}
}

func shortenResponse(w http.ResponseWriter, url string) {
	key := makeKey()
	sites[key] = url

	// pad the output because reasons
	// 1957 + 10 + 33 = 2000
	outputKey := key + randomString(1957)
	fmt.Fprintf(w, "%s/%s", sitePrefix, outputKey)
}

func makeKey() string {
	for {
		key := randomString(10)
		if sites[key] == "" {
			return key
		}
	}
}
