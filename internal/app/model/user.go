package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// User ...
type User struct {
	TG_ID        string `json:"id"`
	TG_FirstName string `json:"first_name"`
	TG_LastName  string `json:"last_name"`
	TG_Username  string `json:"username"`
	TG_PhotoURL  string `json:"photo_url"`
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(u.TG_ID, validation.Required, validation.Length(3, 20)),
	)
}

func (u *User) BeforeCreate() error {
	return nil
}
