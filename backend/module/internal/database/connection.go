package database

import (
    "database/sql"
    "fmt"
    "time"

    "my-project/internal/config"
    _ "github.com/go-sql-driver/mysql"
)

func NewConnection(cfg *config.DBConfig) (*sql.DB, error) {
    db, err := sql.Open("mysql", cfg.DSN())
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }

    // Настройка пула подключений для MySQL
    db.SetMaxOpenConns(cfg.MaxOpenConns)
    db.SetMaxIdleConns(cfg.MaxIdleConns)
    db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)

    // Проверка подключения
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    return db, nil
}