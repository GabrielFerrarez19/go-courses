package services

import (
	"time"

	"librarymvc/internal/users/models"
)

type UserService struct {
	repo models.UserRepository
}

func NewUserService(repo models.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return u.repo.CreateUser(user)
}

func (u UserService) GetUserByID(id int64) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u UserService) ListUSer() ([]*models.User, error) {
	return u.repo.ListUser()
}

func (u UserService) UpdateUsers(user *models.User, id int64) (*models.User, error) {
	user.UpdatedAt = time.Now()

	return u.repo.UpdateUsers(user, id)
}

func (u UserService) DeleteUsers(id int64) error {
	return u.repo.DeleteUsers(id)
}
