package auth

import (
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(user *User) error {
	// Validate user input
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password are required")
	}

	// Check if user already exists
	existingUser, _ := s.repo.FindByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Create user
	return s.repo.CreateUser(user)
}

func (s *Service) Login(username, password string) (*User, error) {
	// Find user by username
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
