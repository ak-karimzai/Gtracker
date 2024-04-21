// Package db provides functionality for interacting with a PostgreSQL database.
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB represents a connection pool to the PostgreSQL database.
type DB struct {
	*pgxpool.Pool
}

// NewPSQL creates a new PostgreSQL connection pool with the provided connection parameters.
// It returns a DB instance and an error if the connection fails.
func NewPSQL(
	Host string,
	Port string,
	Username string,
	DBName string,
	SSLMode string,
	Password string,
) (*DB, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	conn, err := pgxpool.New(timeout,
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			Host, Port, Username, DBName, Password, SSLMode),
	)

	if err != nil {
		return nil, err
	}

	timeout, cancelFunc = context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	err = conn.Ping(timeout)
	return &DB{conn}, err
}
