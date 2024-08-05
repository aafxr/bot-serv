package store

import "github.com/aafxr/chat-bot/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(string) (*model.User, error)
}
