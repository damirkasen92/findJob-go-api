package middleware

type contextKey string

const (
	RequestIDKey contextKey = "request_id" // extracted from request_id.go
	UserIDKey    contextKey = "user_id"
	RoleKey      contextKey = "role"
)
