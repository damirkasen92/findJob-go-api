package auth

import (
	"errors"
	"time"

	"github.com/damir/jobfinder/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint
	Role   model.Role

	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID uint

	jwt.RegisteredClaims
}

type JWTManager struct {
	secret string
}

func NewJWTManager(
	secret string,
) *JWTManager {
	return &JWTManager{
		secret: secret,
	}
}

func (m *JWTManager) GenerateRefreshToken(
	userID uint,
) (string, error) {
	claims := RefreshClaims{
		UserID: userID,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					30 * 24 * time.Hour,
				),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(m.secret),
	)
}

func (m *JWTManager) GenerateToken(
	userID uint,
	role model.Role,
) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					15 * time.Minute, // magic number - it is fine here)
				),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(m.secret),
	)
}

func (m *JWTManager) Parse(
	tokenString string,
) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(
			token *jwt.Token,
		) (interface{}, error) {
			return []byte(m.secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New(
			"invalid token",
		)
	}

	return claims, nil
}

func (m *JWTManager) ParseRefreshToken(
	tokenString string,
) (*RefreshClaims, error) {
	refreshToken, err := jwt.ParseWithClaims(
		tokenString,
		&RefreshClaims{},
		func(
			token *jwt.Token,
		) (interface{}, error) {
			return []byte(m.secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	refreshClaims, ok := refreshToken.Claims.(*RefreshClaims)

	if !ok || !refreshToken.Valid {
		return nil, errors.New(
			"invalid refresh token",
		)
	}

	return refreshClaims, nil
}
