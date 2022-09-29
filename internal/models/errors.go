package models

import "errors"

var (
	ErrDBConn              = errors.New("DB_CONNECTION_ERROR")
	ErrInvalidInput        = errors.New("INVALID_INPUT")
	ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	ErrDBError             = errors.New("DATABASE_ERROR")
)
