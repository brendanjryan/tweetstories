# Tweetstories [![GoDoc](https://godoc.org/github.com/brendanjryan/tweetstories?status.svg)](https://godoc.org/github.com/brendanjryan/tweetstories)
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

A small twitter bot designed to continuously delete all of an account's
tweets which are older than 24 hours.

This application was designed to run indefinitely on Heroku's "free" teir
and will keep itsself alive via a small http server which pings itsself.

### Setup

To use this bot you will need to [create a new twitter app](https://apps.twitter.com/) and
set the following config variables

```bash
# Your twitter consumer key
TWITTER_CONSUMER_KEY

# Your twitter consumer secret
TWITTER_CONSUMER_SECRET

# Your twitter access token
TWITTER_ACCESS_TOKEN

# Your twitter access secret
TWITTER_ACCESS_SECRET

# The http port which the server will listen on
PORT

# The name of your app on heroku
HEROKU_NAME
```

### Development

To build the bot run:
```bash
make
```

To run the bot run:
```bash
make run
```






