package main

import (
	"fmt"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {

	opts := parseOpts()

	urls := make([]string, opts.numUrls)
	for i := 0; i < int(opts.numUrls); i++ {
		urls[i] = fmt.Sprintf("%s/%d", opts.url, int(opts.start)+i)
	}

	targets := make([]vegeta.Target, len(urls))

	for i, url := range urls {
		targets[i] = vegeta.Target{
			Method: opts.method,
			URL:    url,
			Header: opts.headers.Header,
		}
	}

	targeter := vegeta.NewStaticTargeter(targets...)

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	fmt.Printf("Beginning Attack against %d endpoints\n", len(urls))

	for res := range attacker.Attack(targeter, opts.rate, opts.duration) {
		metrics.Add(res)
	}
	metrics.Close()

	out, err := file(opts.outputf, true)
	if err != nil {
		return
	}

	defer out.Close()

	reporter := vegeta.NewTextReporter(&metrics)
	reporter.Report(out)
}
