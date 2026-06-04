package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
)

type UserService interface {
	Register(
		ctx context.Context,
		req dto.RegisterRequest,
	) error

	Login(
		ctx context.Context,
		req dto.LoginRequest,
	) (string, error)
}
