package mw

import (
	"net/http"
	"opsie/pkg/errors"
	"opsie/pkg/logger"
	"opsie/types"
	"time"
)

// logger measures request time, status, and response size.
func HTTPLogger(next types.HandlerFunc) types.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) *errors.Error {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

		err := next(lrw, r)
		duration := time.Since(start)

		statusColor := colorForStatus(lrw.status)
		methodColor := colorForMethod(r.Method)
		sizeColor := "\033[36m"
		reset := "\033[0m"
		bold := "\033[1m"

		fLogger, sLogger := logger.HttpLogger()

		sLogger.Printf("%s%s%-6s%s  %-40s %-22s  %s%3d%s  %8.2fms  %s%dB%s\n",
			bold, methodColor, r.Method, reset,
			r.URL.Path,
			r.RemoteAddr,
			statusColor, lrw.status, reset,
			float64(duration.Microseconds())/1000.0,
			sizeColor, lrw.size, reset,
		)

		fLogger.Printf("%-6s  %-40s  %3d  %8.2fms  %dB\n",
			r.Method,
			r.URL.Path,
			lrw.status,
			float64(duration.Microseconds())/1000.0,
			lrw.size,
		)

		return err
	}
}
// GET     /api/v1/user/owner/count                  <- 192.168.0.201:62493		200		36B  in  2.64ms
type loggingResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(b)
	lrw.size += n
	return n, err
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[32m" // green
	case code >= 300 && code < 400:
		return "\033[36m" // cyan
	case code >= 400 && code < 500:
		return "\033[33m" // yellow
	default:
		return "\033[31m" // red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return "\033[34m" // blue
	case "POST":
		return "\033[32m" // green
	case "PUT":
		return "\033[33m" // yellow
	case "PATCH":
		return "\033[33m" // yellow
	case "DELETE":
		return "\033[31m" // red
	default:
		return "\033[37m" // white
	}
}
