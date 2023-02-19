package helpers

import (
	"math/rand"
	"net/http"
	"time"
)

func RandomDurationMS(maxMS int) time.Duration {
	minMS := 1
	r := (rand.Intn(maxMS-minMS) + minMS)

	return time.Duration(r) * time.Millisecond
}

func Random2xx() int {
	statuses := []int{
		http.StatusOK,
		http.StatusAccepted,
	}

	index := rand.Intn(len(statuses) - 1)

	return statuses[index]
}

func Random4xx() int {
	statuses := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusTooManyRequests,
	}

	index := rand.Intn(len(statuses) - 1)

	return statuses[index]
}

func Random5xx() int {
	statuses := []int{
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	}

	index := rand.Intn(len(statuses) - 1)

	return statuses[index]
}
