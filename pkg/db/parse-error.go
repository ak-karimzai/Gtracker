package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

var (
	ErrNotFound  = errors.New("Not found")
	ErrConflict  = errors.New("Already exist")
	ErrForbidden = errors.New("Forbidden operation")
)

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
