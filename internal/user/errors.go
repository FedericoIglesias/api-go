package user

import (
	"errors"
	"fmt"
)


var ErrFirstNameRequired = errors.New("first name is required")
var ErrLasNameRequired = errors.New("lastname is required")
var ErrEmailRequired = errors.New("email is required")

type ErrNotFound struct{
	ID uint64
}

func (e ErrNotFound)Error()string{
	fmt.Printf("user id: '%d' doesn't exist", e.ID)
	return "nil"
}