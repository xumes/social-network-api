package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is mandatory and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("nick is mandatory and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email is mandatory and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email address")
	}

	if step == "create" && user.Password == "" {
		return errors.New("password is mandatory and cannot be empty")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	return nil
}
