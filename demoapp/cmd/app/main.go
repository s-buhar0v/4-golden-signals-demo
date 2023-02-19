package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/s-buhar0v/demoapp/internal/helpers"
	"github.com/s-buhar0v/demoapp/internal/metrics"
	"github.com/s-buhar0v/demoapp/internal/middleware"
)

func getHTTPRequestsInflightMax() float64 {
	httpRequestsInflightMaxString := os.Getenv("HTTP_REQUESTS_INFLIGHT_MAX")
	httpRequestsInflightMax := 20.0
	if httpRequestsInflightMaxString != "" {
		httpRequestsInflightMax, _ = strconv.ParseFloat(httpRequestsInflightMaxString, 32)
	}

	return httpRequestsInflightMax
}

func main() {
	httpRequestsInflightMax := getHTTPRequestsInflightMax()

	router := chi.NewRouter()
	router.Use(chimiddleware.Logger)
	router.Use(middleware.HTTPMetrics)
	router.Use(middleware.InflightRequests)

	metrics.HttpRequestsInflightMax.WithLabelValues().Set(httpRequestsInflightMax)

	router.Get("/code-2xx", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(helpers.Random2xx())
	})
	router.Get("/code-4xx", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(helpers.Random4xx())
	})
	router.Get("/code-5xx", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(helpers.Random5xx())
	})

	router.Get("/ms-200", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(helpers.RandomDurationMS(200))
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/ms-500", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(helpers.RandomDurationMS(500))
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/ms-1000", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(helpers.RandomDurationMS(1000))
		w.WriteHeader(http.StatusOK)
	})

	router.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
