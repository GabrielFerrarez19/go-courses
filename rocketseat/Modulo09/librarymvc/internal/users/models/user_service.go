package models

type UserService interface {
	CreateUser(user *User) error
	GetUserByID(id int64) (*User, error)
	ListUser() ([]*User, error)
	UpdateUsers(user *User, id int64) (*User, error)
	DeleteUsers(id int64) error
}
