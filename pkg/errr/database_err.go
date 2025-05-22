package errr

import "errors"

var (
	ErrDuplicate   = errors.New("duplicated entry")
	ErrInvalidDate = errors.New("invalid date format")
	ErrNotFound    = errors.New("data not found")
	ErrSQLBuilder  = errors.New("error when building sql statement")
)
