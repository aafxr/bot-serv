package store

import "github.com/aafxr/bot-serv/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(string) (*model.User, error)
}
