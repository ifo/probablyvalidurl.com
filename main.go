package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var sites map[string]string = make(map[string]string)

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const port string = ":3000"
const sitePrefix string = "http://localhost" + port

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

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
		http.Redirect(w, r, sites[r.URL.Path[1:]], http.StatusTemporaryRedirect)
	}
}

func shortenResponse(w http.ResponseWriter, url string) {
	key := makeKey()
	sites[key] = url
	fmt.Fprintf(w, "%s/%s", sitePrefix, key)
}

func makeKey() string {
	for {
		// 1967 + 33 = 2000
		key := randomString(1967)
		if sites[key] == "" {
			return key
		}
	}
}

func randomString(length int) string {
	var res string
	for i := 0; i < length; i++ {
		res += randomChar(randomInt(len(alphabet)))
	}
	return res
}

func randomChar(x int) string {
	return alphabet[x : x+1]
}

func randomInt(x int) int {
	return rand.Intn(x)
}
