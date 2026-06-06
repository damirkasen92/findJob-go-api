package middleware

import (
	"net/http"

	"github.com/damir/jobfinder/internal/httpx"
)

func RequireRole(
	roles ...string,
) func(http.Handler) http.Handler {
	return func(
		next http.Handler,
	) http.Handler {
		return http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {

				role := GetRole(
					r.Context(),
				)

				for _, allowed := range roles {
					if role == allowed {
						next.ServeHTTP(
							w,
							r,
						)

						return
					}
				}

				httpx.Error(
					w,
					http.StatusForbidden,
					"access denied",
				)
			},
		)
	}
}
