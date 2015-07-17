package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var port = flag.Int("port", 3000, "Port to run the server on")

func main() {
	// setup steps
	setupStrings()
	setupRethinkDB()

	// parse all flags
	flag.Parse()

	// setup handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/list", listHandler)

	// serve and catch errors
	log.Printf("Starting server on port %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	body := ""
	if r.Body != nil {
		bdy, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		body = string(bdy)
		if body != "" {
			body, _ = url.QueryUnescape(body)
			// drop "url=" from body
			body = body[4:]
		}
	}

	switch r.URL.Path {
	case "/":
		if body == "" {
			tmpl, _ := template.ParseFiles("templates/index.html")
			tmpl.Execute(w, "paltry")
		} else {
			shortenResponse(w, r, body)
		}
	default:
		// we don't actually care beyond 10 characters
		if len(r.URL.Path) > 10 {
			url := getShortenedUrl(r.URL.Path[1:11]).Url
			if url != "" {
				http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			} else {
				http.NotFound(w, r)
			}
		} else {
			http.NotFound(w, r)
		}
	}
}

func shortenResponse(w http.ResponseWriter, r *http.Request, url string) {
	key, err := saveShortenedUrl(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// pad the output because reasons
	// 1957 + 10 + 33 = 2000
	outputKey := key + randomString(1957)
	fmt.Fprintf(w, "%s://%s/%s", "http", r.Host, outputKey)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	urls, err := getAllUrls()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, urls)
}
