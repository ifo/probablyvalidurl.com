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

const alphabet string = "-_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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
	outputKey := key + unsafeRandomString(1957)
	fmt.Fprintf(w, "%s/%s", sitePrefix, outputKey)
}

func makeKey() string {
	for {
		key := safeRandomString(10)
		if sites[key] == "" {
			return key
		}
	}
}

func safeRandomString(length int) string {
	return randomString(length, true)
}

func unsafeRandomString(length int) string {
	return randomString(length, false)
}

func randomString(length int, safe bool) string {
	var res string
	alphalen := len(alphabet)
	if safe {
		alphalen -= 1
	}
	for i := 0; i < length; i++ {
		res += randomChar(randomInt(alphalen), safe)
	}
	return res
}

func randomChar(x int, safe bool) string {
	if safe {
		return alphabet[1:][x : x+1]
	} else {
		return alphabet[x : x+1]
	}
}

func randomInt(x int) int {
	return rand.Intn(x)
}
