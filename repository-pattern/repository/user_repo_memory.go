package repository

import (
	"errors"
	"repository-pattern/domain"
	"sync"
)

type UserRepositoryMemory struct {
	users map[string]domain.User
	mux   *sync.RWMutex
}

func NewUserRepositoryMemory() UserRepository {
	return &UserRepositoryMemory{
		users: make(map[string]domain.User),
	}
}

func (u *UserRepositoryMemory) Create(user *domain.User) error {
	u.mux.Lock()
	defer u.mux.Unlock()

	if _, exists := u.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	u.users[user.ID] = *user
	return nil
}

func (u *UserRepositoryMemory) Update(user *domain.User) error {
	u.mux.Lock()
	defer u.mux.Unlock()
	if _, exists := u.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	u.users[user.ID] = *user
	return nil
}

func (u *UserRepositoryMemory) FindByID(id string) (*domain.User, error) {
	u.mux.RLock()
	defer u.mux.RUnlock()

	user, exists := u.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (u *UserRepositoryMemory) FindAllUser() ([]domain.User, error) {
	u.mux.RLock()
	defer u.mux.RUnlock()

	users := make([]domain.User, 0, len(u.users))
	for _, user := range u.users {
		users = append(users, user)
	}
	return users, nil
}
