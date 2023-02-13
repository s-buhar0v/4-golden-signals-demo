package metrics

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func HTTPMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srw := NewStatusResponseWriter(w)
		now := time.Now()

		next.ServeHTTP(srw, r)

		elapsedSeocnds := time.Since(now).Seconds()
		pattern := chi.RouteContext(r.Context()).RoutePattern()
		method := chi.RouteContext(r.Context()).RouteMethod
		status := srw.GetStatusString()

		httpRequestsTotal.WithLabelValues(pattern, method, status).Inc()
		httpRequestsDurationHistorgram.WithLabelValues(pattern, method).Observe(elapsedSeocnds)
	})
}
