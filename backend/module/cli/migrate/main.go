package database

import (
    "fmt"
    "log"

    "backend/internal/config"

    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
    m *migrate.Migrate
}

func NewMigrator(db *DB, cfg *config.MigrationConfig) (*Migrator, error) {
    driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to create driver: %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s", cfg.Dir),
        "postgres", 
        driver,
    )
    if err != nil {
        return nil, fmt.Errorf("failed to create migrator: %v", err)
    }

    return &Migrator{m: m}, nil
}