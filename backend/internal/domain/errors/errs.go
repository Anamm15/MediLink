package errs

import "errors"

var (
	ErrNotFound         = errors.New("record not found")
	ErrInternal         = errors.New("internal server error")
	ErrBadParam         = errors.New("bad parameter")
	ErrBadRequest       = errors.New("bad request")
	ErrConflict         = errors.New("record already exists")
	ErrEmailOrPass      = errors.New("invalid email or password")
	ErrOldPassIncorrect = errors.New("password is incorrect")
)
