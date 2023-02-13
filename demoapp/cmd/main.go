package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/s-buhar0v/demoapp/internal/metrics"
)

func randomResponseTimeMS(max int) time.Duration {
	min := 1
	r := (rand.Intn(max-min) + min)

	return time.Duration(r) * time.Millisecond
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(metrics.HTTPMetrics)

	router.Get("/code-200", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/code-400", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	router.Get("/code-500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	router.Get("/ms-200", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(randomResponseTimeMS(200))
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/ms-500", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(randomResponseTimeMS(500))
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/ms-1000", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(randomResponseTimeMS(1000))
		w.WriteHeader(http.StatusOK)
	})

	router.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
