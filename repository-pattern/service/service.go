// service/user_service.go
package service

import (
	"repository-pattern/domain"
	"repository-pattern/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(firstName, lastName, email string) (*domain.User, error) {

	user := &domain.User{
		ID:    (firstName) + (lastName) + (email),
		Name:  firstName + " " + lastName,
		Email: email,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAllUser()
}
