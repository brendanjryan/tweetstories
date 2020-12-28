package main

import (
	"log"
	"os"
	"time"

	"github.com/brendanjryan/tweetstories/server"
	"github.com/getsentry/sentry-go"
)

func main() {
	log.Println("== Running server ==")
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	server.New().Run()

	defer sentry.Flush(2 * time.Second)
	log.Println("== Server terminated ==")
}
