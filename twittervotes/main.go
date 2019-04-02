package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	mgo "gopkg.in/mgo.v2"
)

func init() {
	godotenv.Load()
}

func main() {
	fmt.Printf("Hi there!, %s", os.Getenv("SP_TWITTER_KEY"))
}

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("dialing mongodb: localhost")
	db, err = mgo.Dial("localhost")
	return err
}
func closedb() {
	db.Close()
	log.Println("closed database connection")
}

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}
