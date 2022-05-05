package uerror

import "errors"

func NewError(message string, err error) error {
	return errors.New(message + " " + err.Error())
}
