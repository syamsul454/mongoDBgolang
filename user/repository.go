package user

import (
	"errors"

	"github.com/qiniu/qmgo"
)

type RepositoryUser interface {
	CreateUser(User, error)
}

type repositoryUser struct {
	db *qmgo.Database
}

func NewRepositoryUser(db *qmgo.Database) *repositoryUser {
	return &repositoryUser{db}
}

func (u *repositoryUser) CreateUser(User) (User, error) {
	var data User
	return data, errors.New("dummy eroor")
}
