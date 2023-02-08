package model

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrNotFound            = errors.New("your requested Record is not found")
	ErrConflict            = errors.New("your Record already exist")
	ErrBadParamInput       = errors.New("given Param is not valid")
)
