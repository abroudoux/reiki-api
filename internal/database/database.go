package database

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const DATABASE_PATH = "../../internal/database/database.db"

func InitDatabase() (*sql.DB, error) {
	dir := filepath.Dir(DATABASE_PATH)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", DATABASE_PATH)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTableMessages() error {
	db, err := InitDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS messages (
		id TEXT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		email TEXT,
		message TEXT,
		date TEXT
	);`
	_, err = db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func CreateTableSessions() error {
	db, err := InitDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS sessions (
		id TEXT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		email TEXT,
		date TEXT
	);`
	_, err = db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}