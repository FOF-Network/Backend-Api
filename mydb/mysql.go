package mydb

import "database/sql"

type MySql struct {
	Connection *sql.DB
}

func (db *MySql) GetIDFromToken(Token string) (uint, error) {
	stmt, err := db.Connection.Prepare(`select user_id from user_tokens where token = ?`)
	if err != nil {
		return 0, err
	}

	var id uint
	err = stmt.QueryRow(Token).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *MySql) GetContacts(id uint) ([]*ContactModel, error) {
	stmt, err := db.Connection.Prepare(`select * from contacts where id = ?`)
	if err != nil {
		return 0, err
	}

	contacts := make([]*ContactModel, 0, 10)
	rows, err := stmt.Query(id)
	if err != nil {
		return contacts, err
	}

	for rows.Next() {
		contact := new(ContactModel)
		err = rows.Scan(
			&contact.FirstName, 
			&contact.LastName, 
			&contact.Email, 
			&contact.BirthDay, 
			&contact.Job, 
			&contact.Interests, 
			&contact.CityName)
		if err != nil {
			return contacts, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, rows.Err()
}