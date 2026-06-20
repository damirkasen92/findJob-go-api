package middleware

import (
	"fmt"
	"net/http"

	"github.com/damir/jobfinder/internal/httpx"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				if err, ok := rec.(error); ok {
					httpx.HandleError(w, err)
				} else {
					httpx.HandleError(w, fmt.Errorf("%v", rec))
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
