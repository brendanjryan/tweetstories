package main

import (
	"log"

	"github.com/brendanjryan/tweetstories/server"
)

func main() {
	log.Println("== Running server ==")
	server.New().Run()
	log.Println("== Server terminated ==")
}
