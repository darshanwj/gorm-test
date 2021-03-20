package internal

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func RequestLoggingMiddleware(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Info("request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.EscapedPath()),
				zap.Duration("duration", time.Since(start)))
		}

		return http.HandlerFunc(fn)
	}
}
