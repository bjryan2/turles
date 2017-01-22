# Vegeta ![GoDoc](https://godoc.org/github.com/bjryan2/turles?status.svg)
Turles is a (very) lightweight wrapper around the [vegeta](http://github.com/tsenart/vegeta) loadtesting tool for generating batch load tests over a number of unique urls. It currently supports a *very* constrained featureset and only generates monotonically increasing numeric ids - this may change in the future though :)

## Install
### Pre-compiled executables
Get them [here](http://github.com/bjryan2/turles/releases).

### Source
You need `go` installed and `GOBIN` in your `PATH`. Once that is done, run the
command:
```shell
$ go get -u github.com/bjryan2/turles
```

## Usage manual
```console
Usage: turles url <flags>

flags:
  -duration duration
    	Duration of the test (default 1s)
  -end uint
    	The ending numeric id for request building (default 10)
  -header value
    	Request header
  -method string
    	The HTTP method you wish to use (default "GET")
  -output string
    	Output file (default "stdout")
  -rate uint
    	Requests per second (default 10)
  -start uint
    	The starting numeric id for request building (default 1)
  -url string
    	The URL which you wish to test against

examples:
  turles -url "https://api.foo/users/" -duration=10s -rate 100 
```
