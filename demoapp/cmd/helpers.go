package main

import (
	"math/rand"
	"net/http"
	"time"
)

func randomDurationMS(maxMS int) time.Duration {
	minMS := 1
	r := (rand.Intn(maxMS-minMS) + minMS)

	return time.Duration(r) * time.Millisecond
}

func random4xx() int {
	statuses := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusTooManyRequests,
	}

	index := rand.Intn(len(statuses) - 1)

	return statuses[index]
}

func random5xx() int {
	statuses := []int{
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	}

	index := rand.Intn(len(statuses) - 1)

	return statuses[index]
}
