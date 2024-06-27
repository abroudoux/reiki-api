package database

import "github.com/abroudoux/reiki-api/internal/types"

func GetMessage(id string) (types.Message, error) {
	db, err := InitDatabase()

	if err != nil {
		return types.Message{}, err
	}

	defer db.Close()
	row := db.QueryRow("SELECT id, first_name, last_name, email, message, date FROM messages WHERE id = ?", id)
	var message types.Message
	err = row.Scan(&message.Id, &message.FirstName, &message.LastName, &message.Email, &message.Message, &message.Date)

	if err != nil {
		return types.Message{}, err
	}

	return message, nil
}

func GetMessages() ([]types.Message, error) {
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

func PostMessage(message types.Message) error {
	db, err := InitDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `INSERT INTO messages (id, first_name, last_name, email, message, date) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, message.Id, message.FirstName, message.LastName, message.Email, message.Message, message.Date)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMessage(id string) error {
	db, err := InitDatabase()

	if err != nil {
		return err
	}

	defer db.Close()

	query := `DELETE FROM messages WHERE id = ?`
	_, err = db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}