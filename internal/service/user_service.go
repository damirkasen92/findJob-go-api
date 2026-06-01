package service

import (
	"context"
	"errors"

	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(
	repo repository.UserRepository,
) UserService {

	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(
	ctx context.Context,
	req dto.RegisterRequest,
) error {
	// check if there is any user with the same email
	user, err := s.repo.GetByEmail(
		ctx,
		req.Email,
	)

	if err == nil && user != nil {
		return ErrUserExists
	}

	if err != nil &&
		!errors.Is(
			err,
			gorm.ErrRecordNotFound,
		) {

		return err
	}

	// make hash for a password
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	return s.repo.Create(
		ctx,
		&model.User{
			Email:    req.Email,
			Password: string(hash),
			Role:     "user",
		},
	)
}
