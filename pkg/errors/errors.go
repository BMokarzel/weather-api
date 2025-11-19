package pkg_errors

import "fmt"

type GenericError struct {
	Messgae string
}

func (e *GenericError) Error() string {
	return fmt.Sprintf("Error: %s", e.Messgae)
}

type NotFount struct {
}

func (e *NotFount) Error() string {
	return "Not found"
}

func NewNotFoundError() *NotFount {
	return &NotFount{}
}

func IsNotFound(e error) bool {
	_, ok := e.(*NotFount)
	return ok
}
