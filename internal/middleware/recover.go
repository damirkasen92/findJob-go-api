package middleware

import (
	"net/http"

	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/pkg/logger"
	"go.uber.org/zap"
)

func Recover(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// log the error in the console at the end
			defer func() {
				if err := recover(); err != nil {
					requestID := GetRequestID(
						r.Context(),
					)

					logger.Info("recover failed",
						zap.String("request_id", requestID),
						zap.Any("panic", err),
					)

					httpx.Error(w, http.StatusInternalServerError, "ERROR_INTERNAL_ERROR", "internal server error")
				}
			}()

			// next handler
			next.ServeHTTP(w, r)
		},
	)
}
