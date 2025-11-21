package repository

import (
	"errors"
	"sync"

	"librarymvc/internal/users/models"
)

type UserRepository struct {
	users  map[int64]*models.User
	mu     sync.RWMutex
	nextId int64
}

func NewUserRepository() models.UserRepository {
	return &UserRepository{
		users:  make(map[int64]*models.User),
		nextId: 1,
	}
}

// CreateUser implements models.UserRepository.
func (u *UserRepository) CreateUser(user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user.ID = u.nextId
	u.nextId++

	u.users[user.ID] = user
	return nil
}

// DeleteUsers implements models.UserRepository.
func (u *UserRepository) DeleteUsers(id int64) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	_, exists := u.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(u.users, id)

	return nil
}

// GetUserByID implements models.UserRepository.
func (u *UserRepository) GetUserByID(id int64) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// ListUser implements models.UserRepository.
func (u *UserRepository) ListUser() ([]*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := make([]*models.User, 0, len(u.users))

	for _, user := range u.users {
		users = append(users, user)
	}
	return users, nil
}

// UpdateUsers implements models.UserRepository.
func (u *UserRepository) UpdateUsers(user *models.User, id int64) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	u.users[user.ID] = user

	return user, nil
}
