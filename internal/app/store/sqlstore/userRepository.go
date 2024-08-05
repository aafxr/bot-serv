package sqlstore

import (
	"github.com/aafxr/chat-bot/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (id, first_name, last_name, username, photo) VALUES (?,?,?,?,?)",
		u.TG_ID, u.TG_FirstName, u.TG_LastName, u.TG_Username, u.TG_PhotoURL,
	).Err()
}

// Find ...
func (r *UserRepository) Find(id string) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, first_name, last_name, username, photo FROM users WHERE id=?",
		id,
	).Scan(
		&u.TG_ID,
		&u.TG_FirstName,
		&u.TG_LastName,
		&u.TG_Username,
		&u.TG_PhotoURL,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}
