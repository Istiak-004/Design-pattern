package repository

import "repository-pattern/domain"

type UserRepository interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	FindByID(id string) (*domain.User, error)
	FindAllUser() ([]domain.User, error)
}
