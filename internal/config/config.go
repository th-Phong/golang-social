package config

import (
	"fmt"
	"phongtran/go-social/golang-social/internal/utils"
	"strings"
)

type DatabaseConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

type Config struct {
	ServerAddress string
	Database      DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: fmt.Sprintf(":%s", strings.TrimSpace(utils.GetEnv("SERVER_PORT", "8080"))),
		Database: DatabaseConfig{
			Host:    utils.GetEnv("DB_HOST", "localhost"),
			Port:    utils.GetEnv("DB_PORT", "5432"),
			User:    utils.GetEnv("DB_USER", "postgres"),
			Pass:    utils.GetEnv("DB_PASSWORD", "123456"),
			DBName:  utils.GetEnv("DB_NAME", "db-name"),
			SSLMode: utils.GetEnv("DB_SSLMODE", "disable"),
		},
	}
}

func (c *Config) DNS() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Database.Host, c.Database.Port, c.Database.User,
		c.Database.Pass, c.Database.DBName, c.Database.SSLMode,
	)
}
