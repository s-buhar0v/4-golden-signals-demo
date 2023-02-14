package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/s-buhar0v/demoapp/internal/helpers"
	"github.com/s-buhar0v/demoapp/internal/metrics"
)

func HTTPMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srw := helpers.NewStatusResponseWriter(w)
		now := time.Now()

		next.ServeHTTP(srw, r)

		elapsedSeocnds := time.Since(now).Seconds()
		pattern := chi.RouteContext(r.Context()).RoutePattern()
		method := chi.RouteContext(r.Context()).RouteMethod
		status := srw.GetStatusString()

		metrics.HttpRequestsTotal.WithLabelValues(pattern, method, status).Inc()
		metrics.HttpRequestsDurationHistorgram.WithLabelValues(pattern, method).Observe(elapsedSeocnds)
		metrics.HttpRequestsDurationSummary.WithLabelValues(pattern, method).Observe(elapsedSeocnds)
	})
}
