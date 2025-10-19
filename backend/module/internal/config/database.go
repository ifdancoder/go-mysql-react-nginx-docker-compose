package config

import "fmt"
import "backend/pkg/utils"

type DBConfig struct {
    Host            string
    Port            string
    User            string
    Password        string
    Name            string
    Charset         string
    ParseTime       bool
    Loc             string
    MaxOpenConns    int
    MaxIdleConns    int
    ConnMaxLifetime int
}

// URL для MySQL
func (c *DBConfig) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
        c.User, c.Password, c.Host, c.Port, c.Name, c.Charset, c.ParseTime, c.Loc)
}

// URL для миграций (немного другой формат)
func (c *DBConfig) MigrationDSN() string {
    return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
        c.User, c.Password, c.Host, c.Port, c.Name, c.Charset, c.ParseTime, c.Loc)
}

func NewDBConfig() *DBConfig {
    return &DBConfig{
        Host:            utils.GetEnv("DB_HOST", "localhost"),
        Port:            utils.GetEnv("DB_PORT", "3306"),
        User:            utils.GetEnv("DB_USERNAME", "root"),
        Password:        utils.GetEnv("DB_PASSWORD", ""),
        Name:            utils.GetEnv("DB_DATABASE", "myapp"),
        Charset:         utils.GetEnv("DB_CHARSET", "utf8mb4"),
        ParseTime:       utils.GetEnvBool("DB_PARSE_TIME", true),
        Loc:             utils.GetEnv("DB_LOC", "Local"),
        MaxOpenConns:    utils.GetEnvInt("DB_MAX_OPEN_CONNS", 25),
        MaxIdleConns:    utils.GetEnvInt("DB_MAX_IDLE_CONNS", 25),
        ConnMaxLifetime: utils.GetEnvInt("DB_CONN_MAX_LIFETIME", 5),
    }
}