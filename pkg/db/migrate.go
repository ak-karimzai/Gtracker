package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(migrationUrl string,
	Host string,
	Port string,
	Username string,
	DBName string,
	SSLMode string,
	Password string,
) error {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s", Username, Password, Host, Port, DBName, SSLMode,
	)
	//log.Print(migrationUrl, dbUrl)
	migration, err := migrate.New(migrationUrl, dbUrl)
	if err != nil {
		return fmt.Errorf("cannot create migrate object: %w", err)
	}

	if err = migration.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return fmt.Errorf("failed to up migrate: %w", err)
	}
	return nil
}
