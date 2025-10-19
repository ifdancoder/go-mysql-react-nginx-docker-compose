package database

import (
    "database/sql"
    "fmt"
    "time"

    "backend/internal/config"
)

type DB struct {
    *sql.DB
}

func NewConnection(cfg *config.DBConfig) (*DB, error) {
    db, err := sql.Open("postgres", cfg.URL())
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }

    // Настройка пула подключений
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Проверка подключения
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    return &DB{db}, nil
}

func (db *DB) Close() error {
    return db.DB.Close()
}