package database

import (
	"database/sql"
	"fmt"
	"time"
)

const MAX_OPEN_CONN = 10
const MAX_IDLE_CONN = 5
const CONN_LIFETIME = 10 * time.Minute

type database struct {
	*sql.DB
}

type Configuration struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDB(config *Configuration) (*database, error) {

	connection := fmt.Sprintf("host=%s port=%s password=%s dbname=%s sslmode=disable user=%s",
		config.Host,
		config.Port,
		config.Password,
		config.DBName,
		config.User,
	)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, fmt.Errorf("Connection failed: %w", err)
	}

	db.SetMaxOpenConns(MAX_OPEN_CONN)
	db.SetMaxIdleConns(MAX_IDLE_CONN)
	db.SetConnMaxLifetime(CONN_LIFETIME)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("database is not available to ping")
	}

	return &database{db}, nil 

}

func (db *database) Close() error {
	return db.DB.Close()
}
