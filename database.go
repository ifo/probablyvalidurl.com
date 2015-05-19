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
		log.Fatalln(err.Error())
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

// TODO implement saving
func saveShortenedUrl() error {
	ensureSession()
	return nil
}

// TODO implement getting "shortened" url
func getShortenedUrl() (string, error) {
	ensureSession()
	return "aoeu", nil
}

func ensureSession() {
	if session == nil {
		setupRethinkDB()
	}
}
