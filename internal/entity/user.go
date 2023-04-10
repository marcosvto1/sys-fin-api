package entity

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewUser(name, email, password, confirmPassword, role string) (*User, error) {
	if password != confirmPassword {
		return nil, errors.New("passwords do not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       -1,
		Name:     name,
		Email:    email,
		Password: string(hash),
		Role:     role,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	fmt.Println(err)
	return err == nil
}

func (u *User) ValidateRole(role string) bool {
	switch role {
	case "admin":
		return true
	default:
		return false
	}
}
