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
		return model.ErrUserExists
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
			Role:     model.RoleUser,
		},
	)
}

func (s *userService) Login(
	ctx context.Context,
	req dto.LoginRequest,
) (string, string, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)

	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", "", model.ErrInvalidCredentials
	}

	accessToken, err := s.jwt.GenerateToken(
		user.ID,
		user.Role,
	)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwt.GenerateRefreshToken(
		user.ID,
	)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
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

func (s *userService) Refresh(
	ctx context.Context,
	refreshToken string,
) (string, error) {
	claims, err := s.jwt.ParseRefreshToken(
		refreshToken,
	)

	if err != nil {
		return "", model.ErrInvalidCredentials
	}

	user, err := s.repo.GetByID(
		ctx,
		claims.UserID,
	)

	if err != nil {
		return "", err
	}

	accessToken, err := s.jwt.GenerateToken(
		user.ID,
		user.Role,
	)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
