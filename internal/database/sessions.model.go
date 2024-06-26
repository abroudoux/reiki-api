package database

import "github.com/abroudoux/reiki-api/internal/types"

func PostSession(session types.Session) error {
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

func DeleteSession(id string) error {
	db, err := InitDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `DELETE FROM sessions WHERE id = ?`
	_, err = db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func GetSessions() ([]types.Session, error) {
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