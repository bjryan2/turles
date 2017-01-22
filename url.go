package main

import "fmt"

func genUrls(base string, start uint64, stop uint64) []string {
	urls := make([]string, stop-start)

	for i := start; i < stop; i++ {
		urls[i] = fmt.Sprintf("%s/%d", base, i)
	}

	return urls
}
