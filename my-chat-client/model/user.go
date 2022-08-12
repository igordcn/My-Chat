package model

import (
	"errors"
	"math/rand"
	"time"
)

type User struct {
	id   int
	name string
}

func NewUser(name string) (*User, error) {
	newUser := &User{generateId(), name}
	if err := newUser.validate(); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (user *User) Id() int {
	return user.id
}

func (user *User) Name() string {
	return user.name
}

func generateId() int {
	rand.Seed(time.Now().UnixMilli())
	return rand.Intn(1001)
}

func (user *User) validate() error {
	if user.id < 0 || user.id > 1000 {
		return errors.New("Id is out of bounds. It should be between 0 and 1000")
	}
	if len(user.name) == 0 {
		return errors.New("The name should not be empty!")
	}
	return nil
}
