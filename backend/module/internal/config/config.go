package config

import (
    "backend/pkg/utils"
	"github.com/joho/godotenv"
)

type Config struct {
    DB          *DBConfig
    Migrations  *MigrationConfig
    App         *AppConfig
}

type MigrationConfig struct {
    Dir string
}

type AppConfig struct {
    Env  string
    Port string
}

func New() *Config {
    godotenv.Load()

    return &Config{
        DB: NewDBConfig(),
        Migrations: &MigrationConfig{
            Dir: utils.GetEnv("MIGRATIONS_DIR", "./db/migrations"),
        },
        App: &AppConfig{
            Env:  utils.GetEnv("APP_ENV", "development"),
            Port: utils.GetEnv("APP_PORT", "8080"),
        },
    }
}