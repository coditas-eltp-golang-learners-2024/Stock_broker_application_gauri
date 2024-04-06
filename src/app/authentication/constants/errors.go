package constants

import "errors"

var (
    ErrDatabaseConnection = errors.New("failed to connect to database")
    ErrDatabasePing       = errors.New("failed to ping database")
)
