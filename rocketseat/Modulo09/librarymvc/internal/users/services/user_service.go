package services

import (
	"time"

	"librarymvc/internal/users/models"
)

type UserService struct {
	repo models.UserRepository
}

func NewUserService(repo models.UserRepository) models.UserService {
	return &UserService{}
}

// CreateUser implements models.UserService.
func (u *UserService) CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return u.repo.CreateUser(user)
}

// DeleteUsers implements models.UserService.
func (u *UserService) DeleteUsers(id int64) error {
	return u.repo.DeleteUsers(id)
}

// GetUserByID implements models.UserService.
func (u *UserService) GetUserByID(id int64) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

// ListUser implements models.UserService.
func (u *UserService) ListUser() ([]*models.User, error) {
	return u.repo.ListUser()
}

// UpdateUsers implements models.UserService.
func (u *UserService) UpdateUsers(user *models.User, id int64) (*models.User, error) {
	user.UpdatedAt = time.Now()

	return u.repo.UpdateUsers(user, id)
}
