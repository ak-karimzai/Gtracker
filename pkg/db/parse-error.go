// Package db provides functionality for interacting with the database.
package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

// ErrNotFound is returned when the requested resource is not found in the database.
var ErrNotFound = errors.New("Not found")

// ErrConflict is returned when there is a conflict in the database (e.g., duplicate key).
var ErrConflict = errors.New("Already exist")

// ErrForbidden is returned when the operation is forbidden by the database constraints.
var ErrForbidden = errors.New("Forbidden operation")

// ParseError parses the database error and returns the corresponding application error.
func (db *DB) ParseError(err error) error {
	if err == sql.ErrNoRows || strings.Contains(
		sql.ErrNoRows.Error(), err.Error()) {
		return ErrNotFound
	}

	var pqErr *pgconn.PgError
	if errors.As(err, &pqErr) {
		fmt.Sprint(pqErr.Message, pqErr.Code)
		switch pqErr.Code {
		case "23505":
			err = ErrConflict
		case "23503":
			err = ErrForbidden
		}
	}
	return err
}
