package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type DBConfig interface {
	GetHost() string
	GetPort() string
	GetUsername() string
	GetPassword() string
	GetDatabase() string
}

// Struct that implements this interface and reads configuration from environment
type EnvDBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewEnvDBConfig() *EnvDBConfig {
	return &EnvDBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_DATABASE"),
	}
}

func (c *EnvDBConfig) GetHost() string {
	return c.host
}

func (c *EnvDBConfig) GetPort() string {
	return c.port
}

func (c *EnvDBConfig) GetUsername() string {
	return c.username
}

func (c *EnvDBConfig) GetPassword() string {
	return c.password
}

func (c *EnvDBConfig) GetDatabase() string {
	return c.database
}

func ConnectToDB(config DBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetHost(),
		config.GetPort(),
		config.GetUsername(),
		config.GetPassword(),
		config.GetDatabase(),
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("conenction success")
	}
	return db, nil
}
