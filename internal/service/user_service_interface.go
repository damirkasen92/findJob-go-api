package service

import (
	"context"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
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

	GetByID(
		ctx context.Context,
		id uint,
	) (*model.User, error)
}
