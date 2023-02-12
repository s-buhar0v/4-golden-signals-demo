package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
		},
		[]string{"pattern", "method", "status"},
	)
)

type StatusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (srw *StatusResponseWriter) WriteHeader(status int) {
	srw.status = status
	srw.ResponseWriter.WriteHeader(status)
}

func NewStatusResponseWriter(w http.ResponseWriter) StatusResponseWriter {
	return StatusResponseWriter{w, http.StatusOK}
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusCodeResponseWriter := NewStatusResponseWriter(w)

		next.ServeHTTP(&statusCodeResponseWriter, r)

		pattern := chi.RouteContext(r.Context()).RoutePattern()
		method := chi.RouteContext(r.Context()).RouteMethod
		status := fmt.Sprintf("%d", statusCodeResponseWriter.status)

		fmt.Println(statusCodeResponseWriter.status)

		httpRequestsTotal.WithLabelValues(pattern, method, status).Inc()
	})
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(Metrics)

	router.Get("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/bad_request", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	router.Get("/internal_server_error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	router.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", router)

}
