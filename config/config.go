package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	sentry "github.com/getsentry/sentry-go"
)

const (
	// 6 months in seconds
	defaultMaxAge = float64(86400 * 18)
)

// IFace defines all functilnality provided by the config package.
type IFace interface {
	Logger() *log.Logger
	Twitter() *twitter.Client
	MaxAge() float64
	Port() string
	Name() string
	Sentry() *sentry.Client
}

// Config holds and exposes all configurable and global objects and
// variables.
type Config struct {
	log *log.Logger

	client    *http.Client
	port      string
	name      string
	maxAge    float64
	twitter   *twitter.Client
	sentryDSN string

	sentryOnce   sync.Once
	sentryClient *sentry.Client
}

// New instantiates an instance of Config.
func New() *Config {

	client := Twitter{}.Parse().Client()

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	name := os.Getenv("HEROKU_NAME")
	if name == "" {
		log.Fatalln("!! $HEROKU_NAME not set !!")
	}

	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		log.Fatalln("!! $SENTRY_DSN not set !!")
	}

	var maxAge float64

	maxAgeStr := os.Getenv("TWEETSTORIES_MAX_AGE")

	if maxAgeStr == "" {
		maxAge = defaultMaxAge
	} else {
		maxAge, err := time.ParseDuration(maxAgeStr)
		if err != nil {
			log.Fatalln("unable to parse max age: ", maxAge, err)
		}
	}

	return &Config{
		log:       log.New(os.Stdout, "", log.LstdFlags),
		client:    client,
		port:      addr,
		name:      name,
		maxAge:    maxAge,
		sentryDSN: sentryDSN,
		twitter:   twitter.NewClient(client),
	}
}

// Logger exposes the app-wide logger.
func (c *Config) Logger() *log.Logger {
	return c.log
}

// Twitter exposes a Twitter client.
func (c *Config) Twitter() *twitter.Client {
	return c.twitter
}

// Port exposes the HTTP port of the server.
func (c *Config) Port() string {
	return c.port
}

// Name exposes the name of the app
func (c *Config) Name() string {
	return c.name
}

func (c *Config) Sentry() *sentry.Client {
	c.sentryOnce.Do(func() {
		sentryClient, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: c.sentryDSN,
		},
		)

		if err != nil {
			log.Fatal("error creating sentry client: ", err)
		}

		c.sentryClient = sentryClient
	})

	return c.sentryClient
}

// MaxAge exposes the maximum age of a tweet in terms of delta.
func (c *Config) MaxAge() float64 {
	return c.maxAge
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("!! $PORT not set !!")
	}
	return ":" + port, nil
}

// Twitter holds configuration variables pertaining
// to the Twitter API.
type Twitter struct {
	conKey    string
	conSecret string
	token     string
	secret    string
}

// Parse loads the twitter config variables from the local
// environment.
func (t Twitter) Parse() Twitter {
	t.conKey = os.Getenv("TWITTER_CONSUMER_KEY")
	if t.conKey == "" {
		log.Fatalln("!! $TWITTER_CONSUMER_KEY not set !!")
	}

	t.conSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	if t.conSecret == "" {
		log.Fatalln("!! $TWITTER_CONSUMER_SECRET not set!!")
	}

	t.token = os.Getenv("TWITTER_ACCESS_TOKEN")
	if t.token == "" {
		log.Fatalln("!! $TWITTER_ACCESS_TOKEN not set !!")
	}

	t.secret = os.Getenv("TWITTER_ACCESS_SECRET")
	if t.secret == "" {
		log.Fatalln("!! $TWITTER_ACCESS_SECRET not set !!")
	}

	return t
}

// Client instantiates an HTTP client for
// interacting with the twitter API.
func (t Twitter) Client() *http.Client {
	c := oauth1.NewConfig(t.conKey, t.conSecret)
	tt := oauth1.NewToken(t.token, t.secret)

	return c.Client(oauth1.NoContext, tt)
}
