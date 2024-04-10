package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DB struct {
	*pgxpool.Pool
}

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
