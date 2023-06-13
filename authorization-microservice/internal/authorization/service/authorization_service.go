package service

import (
	"errors"

	"github.com/your-username/authorization-microservice/internal/authorization/model"
	"github.com/your-username/authorization-microservice/internal/authorization/repository"
)

type AuthorizationService struct {
	userRepo *repository.UserRepository
}

func NewAuthorizationService(userRepo *repository.UserRepository) *AuthorizationService {
	return &AuthorizationService{
		userRepo: userRepo,
	}
}

func (s *AuthorizationService) Authenticate(username, password string) (*model.User, error) {
	user, err := s.userRepo.GetUser(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *AuthorizationService) Authorize(user *model.User, role string) error {
	if !user.HasRole(role) {
		return errors.New("unauthorized access")
	}

	return nil
}
