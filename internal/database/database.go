package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/abroudoux/reiki-api/internal/types"
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

func AddSession(session types.Session) error {
	db, err := InitDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `INSERT INTO sessions (id, first_name, last_name, email, date) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(query, session.Id, session.FirstName, session.LastName, session.Email, session.Date)

	if err != nil {
		return err
	}

	return nil
}

func ReturnSessions() ([]types.Session, error) {
	db, err := InitDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query("SELECT id, first_name, last_name, email, date FROM sessions")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var sessions []types.Session

	for rows.Next() {
		var session types.Session
		err := rows.Scan(&session.Id, &session.FirstName, &session.LastName, &session.Email, &session.Date)

		if err != nil {
			return nil, err
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
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

func AddMessage(message types.Message) error {
	db, err := InitDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `INSERT INTO messages (id, first_name, last_name, email, message, date) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(query, message.Id, message.FirstName, message.LastName, message.Email, message.Message, message.Date)

	if err != nil {
		return err
	}

	return nil
}

func ReturnMessages() ([]types.Message, error) {
	db, err := InitDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query("SELECT id, first_name, last_name, email, message, date FROM messages")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var messages []types.Message

	for rows.Next() {
		var message types.Message
		err := rows.Scan(&message.Id, &message.FirstName, &message.LastName, &message.Email, &message.Message, &message.Date)

		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}