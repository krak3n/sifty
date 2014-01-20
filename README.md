# Sifty

[![Build Status](https://travis-ci.org/krak3n/sifty.png?branch=master)](https://travis-ci.org/krak3n/sifty)
[![Coverage Status](https://coveralls.io/repos/krak3n/sifty/badge.png?branch=master)](https://coveralls.io/r/krak3n/sifty?branch=master)

Sifty (name may change) is and will be eventually a Go library for interacting with the Datasift REST API

This is an experimental / learning project so I can learn Go so not recomended for use in production :p

## Usage

    go get github.com/krak3n/sifty

You will get `sifty` bin file in your `$GOPATH/bin` directory.

    ./sifty --help
    Usage of ./sifty:
      -key="": Your datasift api key
      -user="": Your datasift user name

For example:

    ./sifty -user="yourusername" -key="yourapikey"

At the moement this just calls a single API endpoint and returns raw JSON listing the push subscriptions you have.
