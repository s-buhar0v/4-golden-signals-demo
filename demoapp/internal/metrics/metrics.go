package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
		},
		[]string{"pattern", "method", "status"},
	)

	httpRequestsDurationHistorgram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds_historgram",
			Buckets: []float64{
				0.1,  // 100 ms
				0.2,  // 200 ms
				0.25, // 250 ms
				0.5,  // 500 ms
				1,    // 1 s
			},
		},
		[]string{"pattern", "method"},
	)

	httpRequestsDurationSummary = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_duration_seconds_summary",
			Objectives: map[float64]float64{
				0.99: 0.001, // 0.99 +- 0.001
				0.95: 0.01,  // 0.95 +- 0.01
				0.5:  0.05,  // 0.5 +- 0.05
			},
		},
		[]string{"pattern", "method"},
	)
)
