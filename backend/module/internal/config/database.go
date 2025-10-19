package config

import (
	"database/sql"
	"fmt"
	"time"

	"backend/internal/pkg/utils"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func (c *DBConfig) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:            utils.GetEnv("DB_HOST", "localhost"),
		Port:            utils.GetEnv("DB_PORT", "5432"),
		User:            utils.GetEnv("DB_USER", "postgres"),
		Password:        utils.GetEnv("DB_PASSWORD", ""),
		Name:            utils.GetEnv("DB_NAME", "myapp"),
		SSLMode:         utils.GetEnv("DB_SSLMODE", "disable"),
		MaxOpenConns:    utils.GetEnvInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns:    utils.GetEnvInt("DB_MAX_IDLE_CONNS", 25),
		ConnMaxLifetime: time.Duration(utils.GetEnvInt("DB_CONN_MAX_LIFETIME", 5)) * time.Minute,
	}
}