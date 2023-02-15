package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	maxSuccessfulRequests = 8
	maxErrorRequests      = 2
)

var (
	endpoints = []string{
		"/code-200",
		"/code-4xx",
		"/code-5xx",
		"/ms-200",
		"/ms-500",
		"/ms-1000",
	}
	requests = map[string]int{}
)

func randomizeEndpoints() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(endpoints),
		func(i, j int) { endpoints[i], endpoints[j] = endpoints[j], endpoints[i] },
	)
}

func main() {

	for {
		totalRequestsCount := 0

		randomizeEndpoints()

		for _, e := range endpoints {
			requestsToEndpoint := 0
			if e == "/code-200" {
				requestsToEndpoint = rand.Intn(maxSuccessfulRequests)
			} else {
				requestsToEndpoint = rand.Intn(maxErrorRequests)
			}

			requests[e] = requestsToEndpoint
			totalRequestsCount += requestsToEndpoint
		}

		wg := &sync.WaitGroup{}
		wg.Add(totalRequestsCount)

		for endpoint, requestsCount := range requests {
			for i := 0; i < requestsCount; i++ {
				go func(e string) {
					if _, err := http.DefaultClient.Get(
						fmt.Sprintf("http://localhost:8080%s", e),
					); err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}(endpoint)
			}
		}

		wg.Wait()
	}
}
