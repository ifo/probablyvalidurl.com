package main

import (
	"flag"
	r "github.com/dancannon/gorethink"
	"log"
)

var session *r.Session = nil
var db = flag.String("db", "test", "Database name in RethinkDB")
var urlTable = flag.String("table", "urls", "Table name to store URLs in")

func setupRethinkDB() {
	flag.Parse()
	sess, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: *db,
	})
	if err != nil {
		log.Fatalf("RethinkDB Connection error: %s\n", err.Error())
	}
	session = sess

	// ensure named urlTable exists in database
	createTable(*urlTable)
}

func createTable(name string) error {
	ensureSession()
	// TODO return the object in a useable format instead of just the error
	// object, error
	_, err := r.Db(*db).TableCreate(name).Run(session)
	return err
}

type UrlMap struct {
	Id   string `gorethink:"id,omitempty"`
	Url  string `gorethink:"url"`
	Hits int    `gorethink:"hits"`
}

func saveShortenedUrl(url string) (string, error) {
	ensureSession()
	var key string
	for {
		key = randomString(10)
		c, _ := getShortenedUrlCursor(key)
		if c.IsNil() {
			break
		}
	}
	ins := UrlMap{
		Url: url,
		Id:  key,
	}
	_, err := r.Table(urlTable).Insert(ins).Run(session)
	return key, err
}

func getShortenedUrl(key string) UrlMap {
	ensureSession()
	c, err := getShortenedUrlCursor(key)
	if err != nil {
		log.Fatalln(err.Error())
	}
	urlMap, err2 := packageCursorToUrl(c)
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
	err3 := addHitToUrl(urlMap)
	if err3 != nil {
		log.Fatalln(err3.Error())
	}
	return urlMap
}

func getShortenedUrlCursor(key string) (*r.Cursor, error) {
	cursor, err := r.Table(*urlTable).Get(key).Run(session)
	return cursor, err
}

func packageCursorToUrl(c *r.Cursor) (UrlMap, error) {
	rows := []UrlMap{}
	err := c.All(&rows)
	return rows[0], err
}

func addHitToUrl(u UrlMap) error {
	u.Hits += 1
	_, err := r.Table(urlTable).Get(u.Id).Update(u).Run(session)
	return err
}

func getAllUrls() ([]UrlMap, error) {
	ensureSession()
	c, err := r.Table(urlTable).Run(session)
	if err != nil {
		log.Fatalln(err.Error())
	}

	rows := []UrlMap{}
	err2 := c.All(&rows)
	return rows, err2
}

func ensureSession() {
	if session == nil {
		setupRethinkDB()
	}
}
