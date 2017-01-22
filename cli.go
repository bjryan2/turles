package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type opts struct {
	duration time.Duration
	rate     uint64
	start    uint64
	end      uint64
	method   string
	headers  headers
	url      string
	outputf  string
}

func newOpts() *opts {
	return &opts{
		headers: headers{http.Header{}},
	}
}

func parseOpts() *opts {
	fs := flag.NewFlagSet("turles", flag.ExitOnError)

	opts := newOpts()

	fs.StringVar(&opts.outputf, "output", "stdout", "Output file")
	fs.DurationVar(&opts.duration, "duration", 1*time.Second, "Duration of the test")
	fs.Uint64Var(&opts.rate, "rate", 10, "Requests per second")
	fs.Uint64Var(&opts.start, "start", 1, "The starting numeric id for request building")
	fs.Uint64Var(&opts.start, "end", 10, "The ending numeric id for request building")
	fs.StringVar(&opts.url, "url", "", "The URL which you wish to test against")
	fs.StringVar(&opts.method, "method", "GET", "The HTTP method you wish to use")
	fs.Var(&opts.headers, "header", "Request header")

	fs.Parse(os.Args[2:])

	if opts.url == "" {
		log.Fatal("You should provide a url via the -url flag")
	}

	return opts
}
