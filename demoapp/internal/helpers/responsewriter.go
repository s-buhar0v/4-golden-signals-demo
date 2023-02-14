package helpers

import (
	"fmt"
	"net/http"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	status int
}

func NewStatusResponseWriter(w http.ResponseWriter) *StatusResponseWriter {
	return &StatusResponseWriter{w, http.StatusOK}
}

func (srw *StatusResponseWriter) WriteHeader(status int) {
	srw.status = status
	srw.ResponseWriter.WriteHeader(status)
}

func (srw *StatusResponseWriter) GetStatusString() string {
	return fmt.Sprintf("%d", srw.status)
}
