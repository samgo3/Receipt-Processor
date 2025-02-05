package middleware

import (
	"net/http"
	"receipt-processor/internal/utils"
	"time"

	"go.uber.org/zap"
)

// LoggingMiddleware logs the details of each HTTP request and response.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger := utils.GetLogger()
		logger.Info("Request received", zap.String("method", r.Method), zap.String("uri", r.RequestURI))

		defer func() {
			if rec := recover(); rec != nil {
				logger.Error("Request panic", zap.Any("error", rec))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			logger.Info("Request processed", zap.String("method", r.Method), zap.String("uri", r.RequestURI), zap.Duration("duration", time.Since(start)))
		}()
		next.ServeHTTP(w, r)
	})
}
