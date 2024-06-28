package configs

import (
	"os"
)

const (
	// Port
	servicePortEnvPath = "PORT"

	// DB environment paths
	dbHostEnvPath     = "POSTGRES_HOST"
	dbUserEnvPath     = "POSTGRES_USER"
	dbPasswordEnvPath = "POSTGRES_PASSWORD" //nolint:gosec
	dbNameEnvPath     = "POSTGRES_DB"
	dbSchemaEnvPath   = "POSTGRES_SCHEMA"
	dbPortEnvPath     = "POSTGRES_PORT"

	// SMTP environment paths
	smtpHostEnvPath     = "SMTP_HOST"
	smtpPortEnvPath     = "SMTP_PORT"
	smtpUserEnvPath     = "SMTP_USERNAME"
	smtpPasswordEnvPath = "SMTP_PASSWORD"
)

type Configs struct {
	Router *Router
	SMTP   *SMTP
	DB     *DB
}

// SMTP configuration
type SMTP struct {
	Host     string
	Port     string
	Username string
	Password string
}

// DB configuration
type DB struct {
	Host     string
	User     string
	Password string
	DB       string
	Schema   string
	Port     string
}

// Router configuration
type Router struct {
	Port string
}

// GetConfigs initializes and returns configs
func GetConfigs() (*Configs, error) {
	return &Configs{
		Router: &Router{
			Port: os.Getenv(servicePortEnvPath),
		},
		SMTP: &SMTP{
			Host:     os.Getenv(smtpHostEnvPath),
			Port:     os.Getenv(smtpPortEnvPath),
			Username: os.Getenv(smtpUserEnvPath),
			Password: os.Getenv(smtpPasswordEnvPath),
		},
		DB: &DB{
			Host:     os.Getenv(dbHostEnvPath),
			User:     os.Getenv(dbUserEnvPath),
			Password: os.Getenv(dbPasswordEnvPath),
			DB:       os.Getenv(dbNameEnvPath),
			Schema:   os.Getenv(dbSchemaEnvPath),
			Port:     os.Getenv(dbPortEnvPath),
		},
	}, nil
}
