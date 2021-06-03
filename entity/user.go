package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Role string

const (
	USER Role = "USER"
	ADMIN     = "ADMIN"
)

type User struct {
	ID        ID
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	isActive  bool
	Roles     []Role
	Articles  []ID
}

func NewUser (email, password, username string)(*User, error) {
	u := &User{
		ID: NewId(),
		Email: email,
		Username: username,
		CreatedAt: time.Now(),
	}

	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	u.isActive = false
	u.Roles = [0]string{USER}

	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.Email == "" || u.Username == "" || u.Password == "" {
		return ErrInvalidEntity
	}

	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}