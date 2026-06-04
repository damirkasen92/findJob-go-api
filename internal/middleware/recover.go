package middleware

import (
	"log"
	"net/http"

	"github.com/damir/jobfinder/internal/httpx"
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

					log.Printf(
						"request_id=%s panic=%v",
						requestID,
						err,
					)

					httpx.Error(w, http.StatusInternalServerError, "internal server error")
				}
			}()

			// next handler
			next.ServeHTTP(w, r)
		},
	)
}
