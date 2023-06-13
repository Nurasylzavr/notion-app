package repository

import (
	"errors"
	"sync"
	// Import the model package here
)

type UserRepository struct {
	users map[string]model.User
	mutex sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]model.User),
	}
}

func (r *UserRepository) CreateUser(user model.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.users[user.Username]; ok {
		return errors.New("user already exists")
	}

	r.users[user.Username] = user
	return nil
}

func (r *UserRepository) GetUser(username string) (*model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, ok := r.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
