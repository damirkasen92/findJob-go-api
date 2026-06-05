package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/httpx"
)

// "constructor" with DI
type AuthMiddleware struct {
	jwtManager *auth.JWTManager
}

func NewAuthMiddleware(
	jwtManager *auth.JWTManager,
) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: jwtManager,
	}
}

func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Read from headers for Authorization
			authHeader := r.Header.Get(
				"Authorization",
			)

			// if there is no Authorization header
			if authHeader == "" {
				httpx.Error(
					w,
					http.StatusUnauthorized,
					"missing token",
				)

				return
			}

			// trim value from left
			tokenString := strings.TrimPrefix(
				authHeader,
				"Bearer ",
			)

			// parse our trimmed value (token)
			claims, err := m.jwtManager.Parse(
				tokenString,
			)

			if err != nil {
				httpx.Error(
					w,
					http.StatusUnauthorized,
					"invalid token",
				)

				return
			}

			// create a context
			ctx := context.WithValue(
				r.Context(),
				UserIDKey,
				claims.UserID,
			)

			ctx = context.WithValue(
				ctx,
				RoleKey,
				claims.Role,
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		},
	)
}

func GetUserID(
	ctx context.Context,
) uint {
	id, ok := ctx.Value(
		UserIDKey,
	).(uint)

	if !ok {
		return 0
	}

	return id
}

func GetRole(
	ctx context.Context,
) string {
	role, ok := ctx.Value(
		RoleKey,
	).(string)

	if !ok {
		return ""
	}

	return role
}
