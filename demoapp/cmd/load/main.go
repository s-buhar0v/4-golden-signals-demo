package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	endpoints = []string{
		"/code-2xx",
		"/code-4xx",
		"/code-5xx",
		"/ms-200",
		"/ms-500",
		"/ms-1000",
	}
	requests = map[string]int{}
)

func getMaxRequestsCount() (int, int) {
	maxSuccessfulRequests, maxErrorRequests := 15, 5

	maxSuccessfulRequestsString := os.Getenv("HTTP_REQUESTS_SUCCESSFUL_MAX")
	if maxSuccessfulRequestsString != "" {
		maxSuccessfulRequests, _ = strconv.Atoi(maxSuccessfulRequestsString)
	}

	maxErrorRequestsString := os.Getenv("HTTP_REQUESTS_ERROR_MAX")
	if maxErrorRequestsString != "" {
		maxErrorRequests, _ = strconv.Atoi(maxErrorRequestsString)
	}

	return maxSuccessfulRequests, maxErrorRequests

}

func randomizeEndpoints() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(endpoints),
		func(i, j int) { endpoints[i], endpoints[j] = endpoints[j], endpoints[i] },
	)
}

func main() {
	maxSuccessfulRequests, maxErrorRequests := getMaxRequestsCount()

	for {
		totalRequestsCount := 0

		randomizeEndpoints()

		for _, e := range endpoints {
			requestsToEndpoint := 0
			if e == "/code-200" || strings.HasPrefix(e, "/ms-") {
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
