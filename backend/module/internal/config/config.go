package config

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
    return &Config{
        DB: NewDBConfig(),
        Migrations: &MigrationConfig{
            Dir: getEnv("MIGRATIONS_DIR", "./migrations"),
        },
        App: &AppConfig{
            Env:  getEnv("APP_ENV", "development"),
            Port: getEnv("APP_PORT", "8080"),
        },
    }
}