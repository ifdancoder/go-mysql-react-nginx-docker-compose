package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
    m *migrate.Migrate
}

func NewMigrator(db *sql.DB, migrationsPath string) (*Migrator, error) {
    // Получаем абсолютный путь к миграциям
    absPath, err := filepath.Abs(migrationsPath)
    if err != nil {
        return nil, fmt.Errorf("failed to get absolute path: %v", err)
    }

    // Проверяем существование директории
    if _, err := os.Stat(absPath); os.IsNotExist(err) {
        return nil, fmt.Errorf("migrations directory does not exist: %s", absPath)
    }

    // Создаем экземпляр драйвера для PostgreSQL
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to create driver: %v", err)
    }

    // Создаем мигратор
    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s", absPath),
        "postgres", 
        driver,
    )
    if err != nil {
        return nil, fmt.Errorf("failed to create migrator: %v", err)
    }

    return &Migrator{m: m}, nil
}

// Migrate применяет все pending миграции
func (mg *Migrator) Migrate() error {
    if err := mg.m.Up(); err != nil {
        if err == migrate.ErrNoChange {
            log.Println("No new migrations to apply")
            return nil
        }
        return fmt.Errorf("failed to apply migrations: %v", err)
    }
    log.Println("Migrations applied successfully")
    return nil
}

// MigrateTo применяет/откатывает миграции до указанной версии
func (mg *Migrator) MigrateTo(version uint) error {
    if err := mg.m.Migrate(version); err != nil {
        return fmt.Errorf("failed to migrate to version %d: %v", version, err)
    }
    log.Printf("Successfully migrated to version %d", version)
    return nil
}

// Rollback откатывает последнюю миграцию
func (mg *Migrator) Rollback() error {
    if err := mg.m.Steps(-1); err != nil {
        return fmt.Errorf("failed to rollback migration: %v", err)
    }
    log.Println("Rollback completed successfully")
    return nil
}

// RollbackN откатывает N миграций
func (mg *Migrator) RollbackN(steps int) error {
    if err := mg.m.Steps(-steps); err != nil {
        return fmt.Errorf("failed to rollback %d migrations: %v", steps, err)
    }
    log.Printf("Successfully rolled back %d migrations", steps)
    return nil
}

// Force устанавливает версию БД (полезно при corrupted миграциях)
func (mg *Migrator) Force(version int) error {
    if err := mg.m.Force(version); err != nil {
        return fmt.Errorf("failed to force version %d: %v", version, err)
    }
    log.Printf("Successfully forced version to %d", version)
    return nil
}

// Version возвращает текущую версию миграций
func (mg *Migrator) Version() (uint, bool, error) {
    return mg.m.Version()
}

// Close закрывает соединение мигратора
func (mg *Migrator) Close() {
    mg.m.Close()
}