package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/httpx"
	"github.com/damir/jobfinder/internal/model"
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
					"ERROR_MISSING_TOKEN",
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
					"ERROR_INVALID_TOKEN",
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
) model.Role {
	role, ok := ctx.Value(
		RoleKey,
	).(model.Role)

	if !ok {
		return ""
	}

	return role
}
