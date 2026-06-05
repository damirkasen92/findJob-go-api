package service

import (
	"context"
	"errors"

	"github.com/damir/jobfinder/internal/auth"
	"github.com/damir/jobfinder/internal/dto"
	"github.com/damir/jobfinder/internal/model"
	"github.com/damir/jobfinder/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	repo repository.UserRepository
	jwt  *auth.JWTManager
}

func NewUserService(
	repo repository.UserRepository,
	jwt *auth.JWTManager,
) UserService {

	return &userService{
		repo: repo,
		jwt:  jwt,
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

func (s *userService) Login(
	ctx context.Context,
	req dto.LoginRequest,
) (string, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", ErrInvalidCredentials // don't forget to create an error in /project/internal/service/errors.go
	}

	token, err := s.jwt.GenerateToken(
		user.ID,
		user.Role,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetByID(
	ctx context.Context,
	id uint,
) (*model.User, error) {
	return s.repo.GetByID(
		ctx,
		id,
	)
}
