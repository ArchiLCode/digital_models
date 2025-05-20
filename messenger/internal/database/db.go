package database

import (
	"database/sql"
	"fmt"
	"messenger/internal/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(cfg *config.Config) error {
	var err error
	db, err = sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
