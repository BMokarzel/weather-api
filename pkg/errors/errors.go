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
	return "can not find zipcode"
}

func NewNotFoundError() *NotFount {
	return &NotFount{}
}

func IsNotFound(e error) bool {
	_, ok := e.(*NotFount)
	return ok
}

type UnprocessableEntityError struct {
}

func (e *UnprocessableEntityError) Error() string {
	return "invalid zipcode"
}

func NewUnprocessableEntityError() *UnprocessableEntityError {
	return &UnprocessableEntityError{}
}

func IsUnprocessableEntityError(e error) bool {
	_, ok := e.(*UnprocessableEntityError)
	return ok
}

type BadRequestError struct {
}

func (e *BadRequestError) Error() string {
	return "bad request"
}

func NewBadRequestError() *BadRequestError {
	return &BadRequestError{}
}

func IsBadRequestError(e error) bool {
	_, ok := e.(*BadRequestError)
	return ok
}

type InternalServerError struct {
}

func (e *InternalServerError) Error() string {
	return "internal server error"
}

func NewInternalServerError() *InternalServerError {
	return &InternalServerError{}
}

func IsInternalServerError(e error) bool {
	_, ok := e.(*InternalServerError)
	return ok
}
