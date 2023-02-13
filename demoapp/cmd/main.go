package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/s-buhar0v/demoapp/internal/metrics"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(metrics.HTTPMetrics)

	router.Get("/code-200", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/code-4xx", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(random4xx())
	})
	router.Get("/code-5xx", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(random5xx())
	})

	router.Get("/ms-200", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(randomDurationMS(200))
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/ms-500", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(randomDurationMS(500))
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/ms-1000", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(randomDurationMS(1000))
		w.WriteHeader(http.StatusOK)
	})

	router.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
